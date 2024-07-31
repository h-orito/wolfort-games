import Image from 'next/image'
import PrimaryButton from '@/components/button/primary-button'
import Link from 'next/link'
import Head from 'next/head'
import UserInfo from '@/components/pages/user/user-info'
import Footer from '@/components/layout/footer/footer'
import { createInnerClient } from '@/components/graphql/client'
import {
  ChinchiroRoomStatus,
  ChinchiroRoomsQuery,
  QChinchiroRoomsDocument,
  QChinchiroRoomsQuery,
  QChinchiroRoomsQueryVariables,
  SimpleChinchiroRoom
} from '@/lib/generated/graphql'
import { base64ToId } from '@/components/graphql/convert'
import ChinchiroHeader from '@/components/pages/chinchiro/chinchiro-header'

export const getServerSideProps = async () => {
  const client = createInnerClient()
  const { data, error } = await client.query<QChinchiroRoomsQuery>({
    query: QChinchiroRoomsDocument,
    variables: {
      query: {
        statuses: [ChinchiroRoomStatus.Opened]
      } as ChinchiroRoomsQuery
    } as QChinchiroRoomsQueryVariables
  })
  if (error) {
    console.log(error)
    return {
      props: {
        rooms: []
      }
    }
  }
  return {
    props: {
      rooms: data.chinchiroRooms
    }
  }
}

type Props = {
  rooms: SimpleChinchiroRoom[]
}

export default function ChinchiroPage({ rooms }: Props) {
  return (
    <main className='min-h-screen w-full text-center'>
      <Head>
        <title>チンチロ | wolfort games</title>
      </Head>
      <ChinchiroHeader />
      <article className='w-full lg:flex lg:justify-center'>
        <div className='lg:w-[960px] lg:border-x lg:border-gray-300'>
          <div>
            <Image
              src={`/games/images/top.jpg`}
              width={960}
              height={540}
              alt='トップ画像'
            />
          </div>
          <div className='flex-1 p-2 lg:p-4'>
            <div className='my-6'>
              <h1 className='mb-2 text-lg font-bold'>チンチロ</h1>
              <p className='text-xs leading-6'>チンチロです。</p>
            </div>
            <UserInfo />
            <div className='my-6'>
              <h1 className='mb-2 text-lg font-bold'>部屋</h1>
              {rooms.length === 0 ? (
                <p>部屋がありません。</p>
              ) : (
                <table className='mx-auto table-auto border border-gray-500 text-sm'>
                  <thead>
                    <tr className='primary-background'>
                      <th className='border border-gray-500 p-2 text-left'>
                        部屋名
                      </th>
                    </tr>
                  </thead>
                  <tbody>
                    {rooms.map((room: SimpleChinchiroRoom) => {
                      return (
                        <tr key={room.id}>
                          <td className='border p-2'>
                            <Link
                              href={`/chinchiro/rooms/${base64ToId(room.id)}`}
                            >
                              {room.name}
                            </Link>
                          </td>
                        </tr>
                      )
                    })}
                  </tbody>
                </table>
              )}
            </div>
            <div className='my-6'>
              <Link href='/chinchiro/create-room'>
                <PrimaryButton>部屋作成</PrimaryButton>
              </Link>
            </div>
          </div>
        </div>
      </article>
      <Footer />
    </main>
  )
}
