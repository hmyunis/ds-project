import { useState, useEffect, useContext } from 'react'
import { useRouter } from 'next/router'
import { API_URL } from '../../constants'
import { AuthContext } from '../../modules/auth_provider'

const Signup = () => {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const { authenticated } = useContext(AuthContext)
  const router = useRouter()

  useEffect(() => {
    if (authenticated) {
      router.push('/')
    }
  }, [authenticated])

  const submitHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault()

    try {
      const res = await fetch(`${API_URL}/signup`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, email, password }),
      })

      if (res.ok) {
        router.push('/login')
      }
    } catch (err) {
      console.log(err)
    }
  }

  return (
    <div className="m-0 p-0 overflow-hidden w-full min-h-screen flex items-center justify-center">
      {/* Signup Form Container */}
      <div className="relative z-10 max-w-md w-full mx-auto rounded-xl p-8 mt-8 lg:mt-24">
        <form
          onSubmit={submitHandler}
          className="bg-white/10 backdrop-blur-md rounded-2xl shadow-lg p-8 border border-white/20"
        >
          {/* Title */}
          <h2 className="text-3xl font-bold text-white mb-8 text-center">Sign Up</h2>

          {/* Username Input */}
          <div className="relative mb-6">
            <input
              type="text"
              id="username"
              className="peer w-full p-3 bg-transparent outline-none border-0 border-b-2 border-white/30 text-white placeholder-transparent focus:ring-0 focus:border-blue-400 transition-colors"
              placeholder=" "
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
            <label
              htmlFor="username"
              className="absolute left-0 -top-3.5 text-white text-sm transition-all peer-placeholder-shown:text-base peer-placeholder-shown:text-white/50 peer-placeholder-shown:top-2 peer-focus:-top-3.5 peer-focus:text-white peer-focus:text-sm"
            >
              Username
            </label>
          </div>

          {/* Email Input */}
          <div className="relative mb-6">
            <input
              type="email"
              id="email"
              className="peer w-full p-3 bg-transparent outline-none border-0 border-b-2 border-white/30 text-white placeholder-transparent focus:ring-0 focus:border-blue-400 transition-colors"
              placeholder=" "
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
            <label
              htmlFor="email"
              className="absolute left-0 -top-3.5 text-white text-sm transition-all peer-placeholder-shown:text-base peer-placeholder-shown:text-white/50 peer-placeholder-shown:top-2 peer-focus:-top-3.5 peer-focus:text-white peer-focus:text-sm"
            >
              Email Address
            </label>
          </div>

          {/* Password Input */}
          <div className="relative mb-8">
            <input
              type="password"
              id="password"
              className="peer w-full p-3 bg-transparent outline-none border-0 border-b-2 border-white/30 text-white placeholder-transparent focus:ring-0 focus:border-blue-400 transition-colors"
              placeholder=" "
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <label
              htmlFor="password"
              className="absolute left-0 -top-3.5 text-white text-sm transition-all peer-placeholder-shown:text-base peer-placeholder-shown:text-white/50 peer-placeholder-shown:top-2 peer-focus:-top-3.5 peer-focus:text-white peer-focus:text-sm"
            >
              Password
            </label>
          </div>

          {/* Signup Button */}
          <button
            type="submit"
            className="w-full hover:bg-gradient-to-l bg-gradient-to-r border border-cyan-500 from-blue-500 to-cyan-400 text-white py-3 px-6 rounded-lg font-semibold hover:opacity-90 active:scale-95 transition-all duration-500"
          >
            Sign Up
          </button>

          {/* Additional Links */}
          <div className="mt-6 text-center space-y-3">
            <p className="text-white/70 text-sm">
              Already have an account?
              <a href="/login" className="text-cyan-300 hover:text-cyan-400 font-medium">
                {' '}
                Sign in
              </a>
            </p>
          </div>
        </form>
      </div>
    </div>
  )
}

export default Signup
