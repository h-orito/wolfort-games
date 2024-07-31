import { useMutation } from '@apollo/client'
import { useCallback } from 'react'
import { useRouter } from 'next/router'
import { SubmitHandler } from 'react-hook-form'
import { base64ToId } from '@/components/graphql/convert'
import RoomEdit, {
  RoomFormInput
} from '@/components/pages/chinchiro/create-room/room-edit'
import PageHeader from '@/components/pages/page-header'
import Head from 'next/head'
import {
  NewChinchiroRoom,
  RegisterChinchiroRoomDocument,
  RegisterChinchiroRoomMutation,
  RegisterChinchiroRoomMutationVariables
} from '@/lib/generated/graphql'
import ChinchiroHeader from '@/components/pages/chinchiro/chinchiro-header'

export default function CreateGame() {
  const defaultValues = {
    name: '',
    password: ''
  } as RoomFormInput

  const [registerRoom] = useMutation<RegisterChinchiroRoomMutation>(
    RegisterChinchiroRoomDocument,
    {
      onError(error) {
        console.error(error)
      }
    }
  )

  const router = useRouter()
  const onSubmit: SubmitHandler<RoomFormInput> = useCallback(
    async (data) => {
      const { data: resData } = await registerRoom({
        variables: {
          input: {
            name: data.name
          } as NewChinchiroRoom
        } as RegisterChinchiroRoomMutationVariables
      })
      const id = resData?.registerChinchiroRoom?.chinchiroRoom?.id
      if (!id) {
        return
      }
      router.push(`/chinchiro/rooms/${base64ToId(id)}`)
    },
    [registerRoom]
  )

  return (
    <main className='w-full'>
      <Head>
        <title>チンチロ 部屋作成 | wolfort games</title>
      </Head>
      <ChinchiroHeader />
      <article className='lg:flex lg:justify-center'>
        <div className='min-h-screen w-full text-center lg:w-[960px] lg:justify-center lg:border-x lg:border-gray-300'>
          <RoomEdit
            defaultValues={defaultValues}
            onSubmit={onSubmit}
            labelName='作成'
          />
        </div>
      </article>
    </main>
  )
}
