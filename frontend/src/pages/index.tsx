import Image from 'next/image'
import PrimaryButton from '@/components/button/primary-button'
import Link from 'next/link'
import { useState } from 'react'
import Modal from '@/components/modal/modal'
import Term from '@/components/layout/footer/term'
import Policy from '@/components/layout/footer/policy'
import Tip from '@/components/layout/footer/tip'
import Head from 'next/head'
import UserInfo from '@/components/pages/user/user-info'
import Footer from '@/components/layout/footer/footer'

export default function Index() {
  return (
    <main className='min-h-screen w-full lg:flex lg:justify-center'>
      <Head>
        <title>wolfort games</title>
      </Head>
      <article className='flex w-full flex-col text-center lg:w-[960px] lg:justify-center lg:border-x lg:border-gray-300'>
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
            <h1 className='mb-2 text-lg font-bold'>wolfort games</h1>
            <p className='text-xs leading-6'>
              ortが作ったゲームで遊べるサイトです。
            </p>
          </div>
          <UserInfo />
          <div className='my-6'>
            <Link href='/chinchiro'>
              <PrimaryButton>チンチロ</PrimaryButton>
            </Link>
          </div>
        </div>
        <Footer />
      </article>
    </main>
  )
}
