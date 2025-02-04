import { useState, useEffect, useContext } from 'react'
import { API_URL, WEBSOCKET_URL } from '../constants'
import { v4 as uuidv4 } from 'uuid'
import { AuthContext } from '../modules/auth_provider'
import { WebsocketContext } from '../modules/websocket_provider'
import { useRouter } from 'next/router'

const Index = () => {
  const [rooms, setRooms] = useState<{ id: string; name: string }[]>([])
  const [roomName, setRoomName] = useState('')
  const { user } = useContext(AuthContext)
  const { setConn } = useContext(WebsocketContext)
  const router = useRouter()

  const getRooms = async () => {
    try {
      const res = await fetch(`${API_URL}/ws/getRooms`, {
        method: 'GET',
      })
      const data = await res.json()
      if (res.ok) {
        setRooms(data)
      }
    } catch (err) {
      console.log(err)
    }
  }

  useEffect(() => {
    getRooms()
  }, [])

  const submitHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault()
    if (!roomName.trim()) return

    try {
      setRoomName('')
      const res = await fetch(`${API_URL}/ws/createRoom`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          id: uuidv4(),
          name: roomName,
        }),
      })

      if (res.ok) {
        getRooms()
      }
    } catch (err) {
      console.log(err)
    }
  }

  const joinRoom = (roomId: string) => {
    if (user.username) {
      const ws = new WebSocket(
        `${WEBSOCKET_URL}/ws/joinRoom/${roomId}?userId=${user.id}&username=${user.username}`
      )
      if (ws.OPEN) {
        setConn(ws)
        router.push('/app')
      }
    }
  }

  return (
    <div className="m-0 p-0 overflow-hidden w-[600px] mx-auto">
      {/* Main Content */}
      <div className="relative z-10 mx-auto h-screen flex flex-col p-6">
        {/* Header */}
        <div className="my-8">
          <h1 className="text-4xl font-bold text-white text-center">WaveChat Rooms</h1>
          <p className="text-center text-white/70 mt-2">Create or join a chat room</p>
        </div>

        {/* Room Creation Form */}
        <div className="bg-white/10 backdrop-blur-sm rounded-xl p-6 mb-8 border border-white/20">
          <div className="flex gap-4">
            <input
              type="text"
              placeholder="Enter room name..."
              className="flex-1 bg-transparent outline-none text-white placeholder-white/50 border-b-2 border-white/20 focus:border-cyan-400 focus:ring-0 px-2 py-2 transition-colors"
              value={roomName}
              onChange={(e) => setRoomName(e.target.value)}
            />
            <button
              className="bg-cyan-400/90 hover:bg-cyan-400 text-white px-6 py-2 rounded-lg transition-colors flex items-center gap-2"
              onClick={submitHandler}
            >
              <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fillRule="evenodd" d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z" clipRule="evenodd" />
              </svg>
              Create Room
            </button>
          </div>
        </div>

        {/* Available Rooms Section */}
        <div className="bg-white/10 backdrop-blur-sm rounded-xl p-6 border border-white/20 flex-1">
          <h2 className="text-xl font-semibold text-white mb-4">Available Rooms</h2>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            {rooms.map((room) => (
              <div key={room.id} className="bg-white/20 backdrop-blur-sm rounded-lg p-4 border border-white/20 hover:border-cyan-400/50 transition-colors">
                <div className="flex justify-between items-center">
                  <span className="text-white font-medium">{room.name}</span>
                  <button
                    className="bg-cyan-400/90 hover:bg-cyan-400 text-white px-4 py-1.5 rounded-md text-sm transition-colors"
                    onClick={() => joinRoom(room.id)}
                  >
                    Join
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}

export default Index
