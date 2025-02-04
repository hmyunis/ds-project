import { useState, createContext, useEffect, useRef } from 'react'
import { useRouter } from 'next/router'

export type UserInfo = {
  username: string
  id: string
}

export const AuthContext = createContext<{
  authenticated: boolean
  setAuthenticated: (auth: boolean) => void
  user: UserInfo
  setUser: (user: UserInfo) => void
}>({
  authenticated: false,
  setAuthenticated: () => {},
  user: { username: '', id: '' },
  setUser: () => {},
})

const AuthContextProvider = ({ children }: { children: React.ReactNode }) => {
  const [authenticated, setAuthenticated] = useState(false)
  const [user, setUser] = useState<UserInfo>({ username: '', id: '' })

  const router = useRouter()

  useEffect(() => {
    const userInfo = localStorage.getItem('user_info')

    if (!userInfo) {
      if (window.location.pathname != '/signup') {
        router.push('/login')
        return
      }
    } else {
      const user: UserInfo = JSON.parse(userInfo)
      if (user) {
        setUser({
          username: user.username,
          id: user.id,
        })
      }
      setAuthenticated(true)
    }
  }, [authenticated])

  const logout = () => {
    localStorage.removeItem('user_info')
    setAuthenticated(false)
    setUser({ username: '', id: '' })
    router.push('/login')
  }

   // Ref for Vanta.js background
   const vantaRef = useRef<HTMLDivElement>(null)

   useEffect(() => {
     const loadVanta = async () => {
       const script = document.createElement('script')
       script.src = 'https://cdnjs.cloudflare.com/ajax/libs/three.js/r134/three.min.js'
       script.async = true
       script.onload = () => {
         const vantaScript = document.createElement('script')
         vantaScript.src = 'https://cdn.jsdelivr.net/npm/vanta@latest/dist/vanta.waves.min.js'
         vantaScript.async = true
         vantaScript.onload = () => {
           if (vantaRef.current) {
             // @ts-ignore
             window.VANTA.WAVES({
               el: vantaRef.current,
               mouseControls: true,
               touchControls: true,
               gyroControls: false,
               minHeight: 200.0,
               minWidth: 200.0,
               scale: 1.0,
               color: 0x003366, // Dark blue waves
               shininess: 35.0,
               waveHeight: 20.0,
               waveSpeed: 0.5,
               zoom: 0.75,
             })
           }
         }
         document.body.appendChild(vantaScript)
       }
       document.body.appendChild(script)
     }
 
     loadVanta()
   }, [])
 

  return (
    <AuthContext.Provider
      value={{
        authenticated: authenticated,
        setAuthenticated: setAuthenticated,
        user: user,
        setUser: setUser,
      }}
    >
      {/* Vanta.js Background */}
      <div ref={vantaRef} className="absolute top-0 left-0 w-full h-full -z-10"></div>
      
      {children}
      <button
            className='bg-red-600 hover:bg-red-500 fixed bottom-2 left-2 text-white rounded-md py-2 px-4'
            onClick={logout}
          >
            Log out
      </button>
    </AuthContext.Provider>
  )
}

export default AuthContextProvider
