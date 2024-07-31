import Panel from '@/components/panel/panel'
import {
  useChinchiroRoomValue,
  useProgressChinchiroGameValue
} from './chinchiro-room-hook'
import Prologue from './progress-area/prologue'
import Progress from './progress-area/progress'

const ChinchiroProgress = () => {
  const room = useChinchiroRoomValue()
  const progressGame = useProgressChinchiroGameValue()

  if (!room) {
    return <Panel header='進行'>ルームが見つかりません</Panel>
  }

  return (
    <div>
      <Panel header={progressGame == null ? '参加者受付中' : '進行中'}>
        {progressGame == null ? <Prologue /> : <Progress />}
      </Panel>
    </div>
  )
}
export default ChinchiroProgress
