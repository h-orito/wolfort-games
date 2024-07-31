import Modal from '@/components/modal/modal'
import Policy from '@/components/layout/footer/policy'
import Term from '@/components/layout/footer/term'
import Tip from '@/components/layout/footer/tip'
import Link from 'next/link'
import { useState } from 'react'

const Footer = () => {
  const [isOpenTermModal, setIsOpenTermModal] = useState(false)
  const toggleTermModal = (e: any) => {
    if (e.target === e.currentTarget) {
      setIsOpenTermModal(!isOpenTermModal)
    }
  }
  const [isOpenPolicyModal, setIsOpenPolicyModal] = useState(false)
  const togglePolicyModal = (e: any) => {
    if (e.target === e.currentTarget) {
      setIsOpenPolicyModal(!isOpenPolicyModal)
    }
  }
  const [isOpenTipModal, setIsOpenTipModal] = useState(false)
  const toggleTipModal = (e: any) => {
    if (e.target === e.currentTarget) {
      setIsOpenTipModal(!isOpenTipModal)
    }
  }
  return (
    <>
      <footer className='border-t border-gray-300 px-4 py-2 text-xs'>
        <div className='flex justify-center'>
          <Link className='hover:text-blue-500' href='/release-note'>
            更新履歴
          </Link>
          <a
            className='ml-2 cursor-pointer hover:text-blue-500'
            onClick={() => setIsOpenTermModal(true)}
          >
            利用規約
          </a>
          <a
            className='ml-2 cursor-pointer hover:text-blue-500'
            onClick={() => setIsOpenPolicyModal(true)}
          >
            プライバシーポリシー
          </a>
          <a
            className='ml-2 cursor-pointer hover:text-blue-500'
            onClick={() => setIsOpenTipModal(true)}
          >
            投げ銭
          </a>
          <a
            href='https://twitter.com/ort_dev'
            target='_blank'
            className='ml-2 cursor-pointer hover:text-blue-500'
          >
            問い合わせ
          </a>
        </div>
        <div className='flex justify-center'>
          © 2024- ort (
          <a
            href='https://github.com/h-orito/wolfort-games'
            target='_blank'
            className='cursor-pointer hover:text-blue-500'
          >
            GitHub
          </a>
          )
        </div>
      </footer>
      {isOpenTermModal && (
        <Modal header='利用規約' close={toggleTermModal} hideFooter>
          <Term />
        </Modal>
      )}
      {isOpenPolicyModal && (
        <Modal
          header='プライバシーポリシー'
          close={togglePolicyModal}
          hideFooter
        >
          <Policy />
        </Modal>
      )}
      {isOpenTipModal && (
        <Modal header='投げ銭' close={toggleTipModal} hideFooter>
          <Tip />
        </Modal>
      )}
    </>
  )
}
export default Footer
