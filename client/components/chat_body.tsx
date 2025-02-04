import React from 'react'
import { Message } from '../pages/app'

const ChatBody = ({ data }: { data: Array<Message> }) => {
  return (
    <>
      {data.map((message: Message, index: number) => {
        if (message.type === 'self') {
          return (
            <div className='max-w-[60%] ml-auto' key={index}>
              <span className='username text-right block text-cyan-300'>{message.username}</span>
              <div className='bg-gradient-to-l from-cyan-400 to-blue-500 text-white p-3 rounded-xl rounded-tr-none'>
                <p>{message.content}</p>
                <span className='text-xs text-white/70 mt-1 block'>{message.timestamp}</span>
              </div>
            </div>
          )
        } else {
          return (
            <div className='max-w-[60%]' key={index}>
              <span className='username text-white/80'>{message.username}</span>
              <div className='bg-white/20 backdrop-blur-sm text-white p-3 rounded-xl rounded-tl-none'>
                <p>{message.content}</p>
                <span className='text-xs text-white/50 mt-1 block'>{message.timestamp}</span>
              </div>
            </div>
          )
        }
      })}
    </>
  )
}

export default ChatBody
