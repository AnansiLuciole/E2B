import { useRouter } from 'next/router'
import { useSessionContext, useSupabaseClient, useUser } from '@supabase/auth-helpers-react'
import { useEffect, useState } from 'react'

import Spinner from 'components/Spinner'
import { Github } from 'lucide-react'

function SignIn() {
  const supabaseClient = useSupabaseClient()
  const user = useUser()
  const router = useRouter()
  const sessionCtx = useSessionContext()
  const [isSigningIn, setIsSigningIn] = useState(false)

  const isLoading = isSigningIn || sessionCtx.isLoading || !!sessionCtx.session

  useEffect(() => {
    if (user) {
      router.push('/')
    }
  }, [user, router])

  async function signInWithGitHub() {
    try {
      setIsSigningIn(true)
      await supabaseClient.auth.signInWithOAuth({
        provider: 'github',
        options: {
          redirectTo: window.location.href,
          scopes: 'email',
        }
      })
    } catch {
      setIsSigningIn(false)
    }
  }

  return (
    <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-sm space-y-4 items-center justify-center">
        <h2 className="text-center text-2xl font-bold leading-9 tracking-tight text-white">
          Sign in to your account
        </h2>
        <div className="sm:mx-auto sm:w-full sm:max-w-sm justify-center flex">
          <button
            className="flex items-center space-x-2 rounded-md bg-white px-3.5 py-2.5 text-sm font-semibold text-gray-900 shadow-sm hover:bg-gray-100 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white"
            onClick={signInWithGitHub}
            disabled={isLoading}
          >
            {isLoading &&
              <Spinner />
            }
            {!isLoading &&
              <Github size={16} />
            }
            <span>Continue with GitHub</span>
          </button>
        </div>
      </div>
    </div>
  )
}

export default SignIn