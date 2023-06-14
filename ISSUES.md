# Issues
### Clock drift
- https://github.com/firecracker-microvm/firecracker/blob/eb8de3ba1f7cb636d2aaa632fe96b234f3a302e6/FAQ.md#my-guest-wall-clock-is-drifting-how-can-i-fix-it

### Automation
- Separate cluster server and client images
- Stop timestamping cluster disk images

### Disk space
- Diff snapshots

### Session speedup
- `/etc/hosts` lock slowdown?
- Kernel args
- API (Nomad calls, polling, etc.)
- Connecting WS (subscriptions take additional call, etc.)
- Diff snapshots
- Move to ubuntu
- Executing `prisma generate` with 5 cores (smt off) takes 12 sec compared to 4 sec on 2 core (ht on) macbook.

### devbookd update
- Update devbookd in all envs automatically

### API keys, routes
- Move API keys to header
- Fix API rotues so they make sense
- Add automatic refreshing of edit sessions that are currently being snapshotted
- Wait for the end of POST /envs and PATCH /envs/{codeSnippetID} operations - the mechanism is already implemented
- Add access control to the remaining envs routes and for the edit session request
- Fix clock drift fixing so it happens in the first log2(3) seconds
- Add monitoring to the envs routes

### Limit build/update env CPU and memory

### Limit session length - is a session runnign 48 hours really something we want to allow? These sessions also accumulate over time

### Add balooning

### Delete files after a failed build/update env
- Stop possible running containers and delete docker image

### fc-envs build scripts are sometimes cached!

### Rebuild only the changed templates on push

### CLI
- fix error whe disconnecting from `connect`
- `template is a required field` error when trying to `use [envID]`
- Add a warning about scanning local fs when the ls command takes too long
- CLI program name should be obvious from the NPM package name
- Fix CI/CD deployment (pnpm, SDK)
- Integrate with environmentID in `guide.json`
- Add published/unpublished label to `env list` output
- Add disclaimer when the env is unpublished to `env push`
- Change `/` root to `/code` root in 
  - Handle deleting files that were removed in local filesystem - how can we recognize that the file was deleted? Saving what was pushed last and only manipulating with these files? (lock) If the file changed since then we should probably ignore it (save name+hash that we already calculate for everyting together)
- Allow custom string for template (--select option for the current list?)
- Fix size and dependencies
- Change template option description
- Finish CLI README
- Warn about having multiple configs for the same environment
- Put template field from dbk.toml in a separate [] category so it is not confused with id because they can look the same
- Add default "empty" template that is used if you don't define anything
- Ignore "files" directory when pushing if it is not present
- Connect with id may be creating toml?
#### Feedback
> How can I create the DEVBOOK_KEY? It seems I can’t create resources without it
> I haven’t used the dashboard before and I presumed I might have needed an invite to a team or sth of the sort.
> Can I change my email? (dashboard)
> Should I store the .dbk files in the repo?
> I just realized the only way I can edit a file in an environment is using vim. I’m not familiar with vim, and editing files might be tedious. Is there a way I can connect to an environment using VS Code Remote - SSH extension?
> Should the ./files directory mimic the file system on the environment? i.e. have a code/prisma/migrations directory and other files?
- It should contain just the files you want to be there. Right now the command finds all the files and uploads them — if overwrites the old files but if there were some files or directories already it will not delete them — for now you can do that via the “connect” and rm command in the env.
> I’ve only connected to the environment to create the initial migration and install dependencies. I might also set up a seed script/ file to seed the database with data. (that might be an extra step)
> Can I update an environment once I’ve published it?
> Would updating one environment and publishing it cascade to the other environments that depend on it?


### Envs system
- "ad-hoc" no provision if hash match
- git based workflow
- unlimited forking
- minimal overhead with diff snapshots
The system would work similarly to git - you have env config inside a repository and every time you push a new environment will be forked from the old environment - its whole id will be <envID>-<changeToEnvHash> where change to env has represents the changes made by push. The has will be also saved to the config identifying environment so we know what environment to fork next time we push.

We would need to keep all environments forks because wa want to be able to access it like in a git -- when you checkout old commit you can get the old environment. This will also help with the diff snapshots - if you start the environment the actual snpashot would be made by composing all the snapshots of the fork's "parents", reducing space needed.

We need to use rootfs CoW **and** diff snapshots for this, because otherwise each change and push would need at least [amount of FC RAM] data of extra space.

This will allow you to create "common base" environments just by copying environent that you want to base the new env on and doing push.

Even setup from terminal/installed packages can be cached this way if we just modify the env cache either by diggesting all the input you entered into terminal or naively by just diggesting some randomized value everytime you connect to the edit session (maybe there are no edit sessions now?) by terminal, creating a new hash. Maybe the fork will happen everytime jsut before you connect to edit session so we can keep the environments immutable.

There will be no "templates", just other envs you want to base the current env on and you don't have to explicitly say "base this new env on existing env with <envID>", you just start using the old env and make changes to it and a new environment with the changes will be automatically created.

The template field that each environment has will be used for composing the diff snapshots and can be also used for constructing the "tree of 

- How to turn this into serverless?

# envs
- How to manage envs with git workflow?
- Add memory swap to Rust - OoM when compiling bigger programs
  - https://www.digitalocean.com/community/tutorials/how-to-add-swap-space-on-ubuntu-22-04
- Should we use ubuntu 22 for envs?
- Does ubuntu need the newer kernel?
  - Using driver with multiple kernels
- Cloud Hypervisor - https://github.com/cloud-hypervisor/cloud-hypervisor