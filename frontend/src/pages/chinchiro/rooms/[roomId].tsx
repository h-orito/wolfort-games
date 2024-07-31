import { createInnerClient } from '@/components/graphql/client'
import { idToBase64 } from '@/components/graphql/convert'
import ChinchiroHeader from '@/components/pages/chinchiro/chinchiro-header'
import ChinchiroChatArea from '@/components/pages/chinchiro/room/chat-area'
import {
  useChinchiroGamesInit,
  useChinchiroRoom,
  useMyChinchiroRoomParticipantInit
} from '@/components/pages/chinchiro/room/chinchiro-room-hook'
import ChinchiroParticipants from '@/components/pages/chinchiro/room/participants'
import ChinchiroProgress from '@/components/pages/chinchiro/room/progress-area'
import { useMyPlayer } from '@/components/pages/user/player-hooks'
import {
  ChinchiroRoom,
  QChinchiroRoomDocument,
  QChinchiroRoomQuery,
  QChinchiroRoomQueryVariables
} from '@/lib/generated/graphql'
import Head from 'next/head'

export const getServerSideProps = async (context: any) => {
  const { roomId } = context.params
  const client = createInnerClient()
  const roomStringId = idToBase64(roomId, 'ChinchiroRoom')
  const { data: roomdata } = await client.query<QChinchiroRoomQuery>({
    query: QChinchiroRoomDocument,
    variables: { roomId: roomStringId } as QChinchiroRoomQueryVariables
  })
  const room = roomdata.chinchiroRoom as ChinchiroRoom
  return {
    props: {
      room: room
    }
  }
}

type Props = {
  room: ChinchiroRoom
}

const ChinchiroRoomPage = ({ room }: Props) => {
  useChinchiroRoom(room)
  useMyPlayer()
  useMyChinchiroRoomParticipantInit(room.id)
  useChinchiroGamesInit(room.id)

  return (
    <main className='w-full'>
      <Head>
        <title>{room.name} | チンチロ | wolfort games</title>
      </Head>
      <ChinchiroHeader />
      <article className='w-full text-sm lg:flex lg:justify-center'>
        <div className='p-2 lg:w-[1280px] lg:p-4'>
          <div className='mb-4 w-full text-center'>
            <h1>{room.name}</h1>
          </div>
          <div className='w-full md:flex md:gap-6'>
            <div className='max-w-full md:max-w-[360px]'>
              <div>
                <ChinchiroParticipants />
              </div>
              <div className='mt-4'>
                <ChinchiroProgress />
              </div>
            </div>
            <div className='mt-4 w-full md:mt-0 md:flex-1'>
              <ChinchiroChatArea />
            </div>
          </div>
        </div>
      </article>
    </main>
  )
}
export default ChinchiroRoomPage
