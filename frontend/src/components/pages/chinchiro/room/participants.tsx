import Panel from '@/components/panel/panel'
import {
  useChinchiroRoomValue,
  useMyChinchiroRoomParticipantValue
} from './chinchiro-room-hook'
import {
  ChinchiroRoom,
  ChinchiroRoomParticipant
} from '@/lib/generated/graphql'

const ChinchiroParticipants = () => {
  const room = useChinchiroRoomValue()
  const myRoomParticipant = useMyChinchiroRoomParticipantValue()

  if (!room) {
    return <Panel header='参加者'>ルームが見つかりません</Panel>
  } else if (room.participants.length === 0) {
    return <Panel header='参加者'>参加者がいません</Panel>
  }

  return (
    <Panel header='参加者' detailsClassName=''>
      {room.participants.map((participant) => (
        <ChinchiroParticipant
          key={participant.id}
          room={room}
          participant={participant}
          isMyself={myRoomParticipant?.id === participant.id}
        />
      ))}
    </Panel>
  )
}
export default ChinchiroParticipants

type ChinchiroRoomParticipantProps = {
  room: ChinchiroRoom
  participant: ChinchiroRoomParticipant
  isMyself: boolean
}

const ChinchiroParticipant = ({
  room,
  participant,
  isMyself
}: ChinchiroRoomParticipantProps) => {
  const isRoomMaster = room.roomMasters.some(
    (rm) => rm.player.id === participant.player.id
  )
  return (
    <div key={participant.id} className='flex items-center px-4 py-2'>
      <div className={`flex-1 ${isMyself ? 'font-bold' : ''}`}>
        {participant.name}
      </div>
      <div>
        {isRoomMaster && (
          <span className='rounded-full bg-yellow-200 px-2 py-1 text-xs text-yellow-800'>
            部屋主
          </span>
        )}
      </div>
    </div>
  )
}
