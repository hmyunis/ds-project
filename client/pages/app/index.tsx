import React, { useState, useRef, useContext, useEffect } from 'react'
import ChatBody from '../../components/chat_body'
import { WebsocketContext } from '../../modules/websocket_provider'
import { useRouter } from 'next/router'
import { API_URL } from '../../constants'
import autosize from 'autosize'
import { AuthContext } from '../../modules/auth_provider'

export type Message = {
  content: string
  client_id: string
  username: string
  room_id: string
  timestamp: string
  type: 'recv' | 'self'
}

const index = () => {
  const [messages, setMessages] = useState<Array<Message>>([])
  const textarea = useRef<HTMLInputElement>(null)
  const { conn } = useContext(WebsocketContext)
  const [users, setUsers] = useState<Array<{ username: string }>>([])
  const { user } = useContext(AuthContext)

  const router = useRouter()

  useEffect(() => {
    if (!conn) {
      router.push('/')
      return
    }

    const roomId = conn.url.split('/')[5]

    const fetchMessages = async () => {
      try {
        const res = await fetch(`${API_URL}/ws/getMessages/${roomId}`)
        const data = await res.json()
        setMessages(
          data.map((msg: any) => ({
            ...msg,
            type: msg.username === user?.username ? 'self' : 'recv',
          }))
        )
      } catch (error) {
        console.error(error)
      }
    }

    const fetchUsers = async () => {
      try {
        const res = await fetch(`${API_URL}/ws/getClients/${roomId}`)
        const data = await res.json()
        setUsers(data)
      } catch (error) {
        console.error(error)
      }
    }

    fetchMessages()
    fetchUsers()
  }, [conn, router, user])

  useEffect(() => {
    if (textarea.current) autosize(textarea.current)

    if (!conn) {
      router.push('/')
      return
    }

    conn.onmessage = (message) => {
      const msg: Message = JSON.parse(message.data)

      if (msg.content === 'A new user has joined the room') {
        setUsers((prev) => [...prev, { username: msg.username }])
      } else if (msg.content === 'user left the chat') {
        setUsers((prev) => prev.filter((u) => u.username !== msg.username))
      }

      setMessages((prev) => [
        ...prev,
        { ...msg, type: msg.username === user?.username ? 'self' : 'recv' },
      ])
    }
  }, [conn, user])

  const sendMessage = () => {
    if (!textarea.current?.value || !conn) return
    conn.send(textarea.current.value)
    textarea.current.value = ''
  }

  return (
    <div className='relative z-10 w-2/3 mx-auto h-screen flex flex-col'>
      {/* Chat Header */}
      <div className='bg-white/10 backdrop-blur-sm py-4 px-6 rounded-t-xl border-b border-white/20'>
        <div className='flex items-center space-x-3'>
          <div className='h-10 w-10 rounded-full bg-cyan-400'></div>
          <div>
            <h2 className='text-white font-semibold'>Team Chat</h2>
            <p className='text-sm text-cyan-200'>{users?.length} Online</p>
          </div>
        </div>
      </div>

      {/* Chat Messages */}
      <div className='chat-container flex-1 overflow-y-auto p-4 space-y-4'>
        <ChatBody data={messages} />
      </div>

      {/* Message Composer */}
      <div className='bg-white/10 backdrop-blur-sm p-4 border-t border-white/20'>
        <div className='flex space-x-2'>
          <input
            type='text'
            placeholder='Type your message...'
            className='flex-1 bg-transparent outline-none text-white placeholder-white/50 border-b-2 border-white/20 focus:border-cyan-400 focus:ring-0 px-2 py-2 transition-colors'
            ref={textarea}
          />
          <button
            className='bg-cyan-400/90 hover:bg-cyan-400 text-white p-2 rounded-lg transition-colors'
            onClick={sendMessage}
          >
            <svg xmlns='http://www.w3.org/2000/svg' className='h-6 w-6' fill='none' viewBox='0 0 24 24' stroke='currentColor'>
              <path strokeLinecap='round' strokeLinejoin='round' strokeWidth='2' d='M12 19l9 2-9-18-9 18 9-2zm0 0v-8' />
            </svg>
          </button>
        </div>
      </div>
    </div>
  )
}

export default index
