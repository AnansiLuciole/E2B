variable "gcp_zone" {
  type = string
}

variable "client_cluster_size" {
  type = number
  default = 0
}

variable "session_proxy_job_index" {
  type = number
  default = 0
}

job "client-proxy" {
  datacenters = [var.gcp_zone]

  group "client-proxy" {
    network {
      port "health" {
        static = 3001
      }
      port "session" {
        static = 3002
      }
    }

    service {
      name = "client-proxy"
      port = "session"
      meta {
        # We force recalculation of the template by redeploying the job when we change the cluster size
        SessionProxyJobIndex = var.session_proxy_job_index
        ClientClusterSize = var.client_cluster_size
      }
    }

    task "nginx" {
      driver = "docker"

      config {
        image = "nginx"
        ports = ["health", "session"]

        volumes = [
          "local:/etc/nginx/conf.d",
        ]
      }

      template {
        left_delimiter  = "[["
        right_delimiter = "]]"
        destination   = "local/load-balancer.conf"
        change_mode   = "signal"
        change_signal = "SIGHUP"
        data            = <<EOF
[[ range service "session-proxy" ]]
server {
  listen 3002;
  server_name ~^\w+_[[ index .ServiceMeta "Client" ]]\.ondevbook\.com$;
  proxy_set_header Host $host;
  proxy_set_header X-Real-IP $remote_addr;
  location / {
    proxy_pass $scheme://[[ .Address ]]:[[ .Port ]]$request_uri;
  }
}
[[ end ]]
server {
  listen 3001;
  location /__health {
    access_log off;
    return 200 '{"status":"UP"}';
    add_header 'Content-Type' 'application/json';
  }
}
EOF
      }
    }
  }
}
