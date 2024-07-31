import { useMyPlayerValue } from '@/components/pages/user/player-hooks'
import {
  useChinchiroRoomValue,
  useIsRoomMaster,
  useMyChinchiroRoomParticipantValue
} from '../chinchiro-room-hook'
import InputText from '@/components/form/input-text'
import { SubmitHandler, useForm } from 'react-hook-form'
import {
  LeaveChinchiroRoom,
  MleaveChinchiroRoomDocument,
  MleaveChinchiroRoomMutation,
  MleaveChinchiroRoomMutationVariables,
  NewChinchiroGame,
  NewChinchiroRoomParticipant,
  ParticipateChinchiroRoomDocument,
  ParticipateChinchiroRoomMutation,
  ParticipateChinchiroRoomMutationVariables,
  RegisterChinchiroGameDocument,
  RegisterChinchiroGameMutation,
  RegisterChinchiroGameMutationVariables
} from '@/lib/generated/graphql'
import SubmitButton from '@/components/button/submit-button'
import { useMutation } from '@apollo/client'
import { useRouter } from 'next/router'
import { useCallback, useMemo } from 'react'

const Prologue = () => {
  const room = useChinchiroRoomValue()
  const myRoomParticipant = useMyChinchiroRoomParticipantValue()
  const myPlayer = useMyPlayerValue()
  const isRoomMaster = room && useIsRoomMaster(room.id)

  const canParticipate = useMemo(() => {
    return myRoomParticipant == null && myPlayer != null
  }, [myRoomParticipant, myPlayer])
  const canLeave = useMemo(() => {
    return myRoomParticipant != null
  }, [myRoomParticipant])

  return (
    <div>
      <p>
        現在ゲーム開始前です。
        <br />
        2人以上集まれば部屋主がゲームを開始することができます。
      </p>
      {canParticipate && <ParticipateArea />}
      {canLeave && <LeaveArea />}
      {isRoomMaster && <StartGameArea />}
    </div>
  )
}

export default Prologue

interface ParticipateInput {
  name: string
}

const ParticipateArea = () => {
  const room = useChinchiroRoomValue()!
  const player = useMyPlayerValue()!
  const { control, formState, handleSubmit } = useForm<ParticipateInput>({
    defaultValues: {
      name: player.name
    }
  })
  const canSubmit: boolean = !formState.isSubmitting

  const [participate] = useMutation<ParticipateChinchiroRoomMutation>(
    ParticipateChinchiroRoomDocument,
    {
      onError(error) {
        console.error(error)
      }
    }
  )

  const router = useRouter()
  const onSubmit: SubmitHandler<ParticipateInput> = useCallback(
    async (data) => {
      const { data: resData } = await participate({
        variables: {
          input: {
            roomId: room.id,
            name: data.name,
            password: '' // TODO
          } as NewChinchiroRoomParticipant
        } as ParticipateChinchiroRoomMutationVariables
      })
      router.reload()
    },
    [participate]
  )

  return (
    <div>
      <hr className='my-4 border-gray-300' />
      <form onSubmit={handleSubmit(onSubmit)}>
        <strong>ニックネーム</strong>
        <InputText
          name='name'
          control={control}
          rules={{
            required: '必須です',
            maxLength: {
              value: 10,
              message: `10文字以内で入力してください`
            }
          }}
        />
        <div className='mt-4'>
          <SubmitButton disabled={!canSubmit} label='参加する' />
        </div>
      </form>
    </div>
  )
}

interface LeaveInput {}

const LeaveArea = () => {
  const room = useChinchiroRoomValue()!
  const myRoomParticipant = useMyChinchiroRoomParticipantValue()!
  const { handleSubmit } = useForm<LeaveInput>()

  const [leave] = useMutation<MleaveChinchiroRoomMutation>(
    MleaveChinchiroRoomDocument,
    {
      onError(error) {
        console.error(error)
      }
    }
  )

  const router = useRouter()
  const onSubmit: SubmitHandler<LeaveInput> = useCallback(
    async (data) => {
      const { data: resData } = await leave({
        variables: {
          input: {
            roomId: room.id
          } as LeaveChinchiroRoom
        } as MleaveChinchiroRoomMutationVariables
      })
      router.reload()
    },
    [leave]
  )

  return (
    <div>
      <hr className='my-4 border-gray-300' />
      <form onSubmit={handleSubmit(onSubmit)}>
        <div>
          <SubmitButton label='退出する' />
        </div>
      </form>
    </div>
  )
}

interface StartGameInput {}

const StartGameArea = () => {
  const room = useChinchiroRoomValue()!
  const canStartGame = useMemo(() => {
    return room?.participants.length >= 2
  }, [room])

  const { handleSubmit } = useForm<StartGameInput>()

  const [startGame] = useMutation<RegisterChinchiroGameMutation>(
    RegisterChinchiroGameDocument,
    {
      onError(error) {
        console.error(error)
      }
    }
  )

  const onSubmit: SubmitHandler<StartGameInput> = useCallback(
    async (data) => {
      const { data: resData } = await startGame({
        variables: {
          input: {
            roomId: room.id
          } as NewChinchiroGame
        } as RegisterChinchiroGameMutationVariables
      })
    },
    [startGame]
  )

  return (
    <div>
      <hr className='my-4 border-gray-300' />
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className=''>
          <SubmitButton disabled={!canStartGame} label='ゲームを開始する' />
        </div>
      </form>
    </div>
  )
}
