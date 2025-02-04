import { useState, useContext, useEffect } from 'react'
import { API_URL } from '../../constants'
import { useRouter } from 'next/router'
import { AuthContext, UserInfo } from '../../modules/auth_provider'

const Login = () => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const { authenticated } = useContext(AuthContext)
  const router = useRouter()

  useEffect(() => {
    if (authenticated) {
      router.push('/')
      return
    }
  }, [authenticated])

  const submitHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault()

    try {
      const res = await fetch(`${API_URL}/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      })

      const data = await res.json()
      if (res.ok) {
        const user: UserInfo = {
          username: data.username,
          id: data.id,
        }

        localStorage.setItem('user_info', JSON.stringify(user))
        return router.push('/')
      }
    } catch (err) {
      console.log(err)
    }
  }

  return (
    <div className="relative flex items-center justify-center min-h-screen w-full">
      {/* Login Form */}
      <form className="relative z-10 bg-white/10 backdrop-blur-sm rounded-xl p-8 w-[350px] shadow-md border border-white/20">
        <h1 className="text-3xl font-bold text-white text-center">Welcome!</h1>

        <input
          placeholder="Email"
          className="w-full bg-transparent outline-none text-white placeholder-white/50 border-b-2 border-white/20 focus:border-cyan-400 focus:ring-0 px-2 py-3 mt-6 transition-colors"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />

        <input
          type="password"
          placeholder="Password"
          className="w-full bg-transparent outline-none text-white placeholder-white/50 border-b-2 border-white/20 focus:border-cyan-400 focus:ring-0 px-2 py-3 mt-4 transition-colors"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />

        <button
          className="w-full bg-cyan-400/90 hover:bg-cyan-400 text-white font-bold py-3 rounded-lg mt-6 transition-colors"
          type="submit"
          onClick={submitHandler}
        >
          Login
        </button>

        <p className="text-center text-white/70 mt-4">
          Don't have an account? <a href="/signup" className="text-cyan-300 hover:text-cyan-400">Sign up!</a>
        </p>
      </form>
    </div>
  )
}

export default Login
