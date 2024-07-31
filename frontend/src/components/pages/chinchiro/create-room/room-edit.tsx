import SubmitButton from '@/components/button/submit-button'
import InputText from '@/components/form/input-text'
import { useState } from 'react'
import { useForm, SubmitHandler } from 'react-hook-form'
import { QuestionMarkCircleIcon } from '@heroicons/react/24/outline'
import Modal from '@/components/modal/modal'

type Props = {
  defaultValues: RoomFormInput
  onSubmit: SubmitHandler<RoomFormInput>
  labelName: string
}

export interface RoomFormInput {
  name: string
  password: string
}

export default function RoomEdit(props: Props) {
  const { control, formState, handleSubmit } = useForm<RoomFormInput>({
    defaultValues: props.defaultValues
  })
  const canSubmit: boolean = !formState.isSubmitting

  return (
    <div>
      <div className='p-4'>
        <div className='flex justify-center'>
          <p className='notification-background notification-text my-2 p-4 text-xs'>
            <span className='text-red-500'>*&nbsp;</span>
            がついている項目は必須です。
            <br />
            また、作成後に変更することができます。
          </p>
        </div>
        <form onSubmit={handleSubmit(props.onSubmit)}>
          <div className='my-4'>
            <FormLabel label='部屋名' required />
            <InputText
              name='name'
              control={control}
              rules={{
                required: '必須です',
                maxLength: {
                  value: 50,
                  message: `50文字以内で入力してください`
                }
              }}
            />
          </div>
          <hr />
          <div className='my-4'>
            <FormLabel label='参加パスワード'>
              設定されている場合、このパスワードを入力しないと参加することができません。
              <br />
              閲覧はパスワードを知らなくても可能です。
            </FormLabel>
            <div className='flex justify-center'>
              <p className='notification-background notification-text my-2 p-4 text-xs'>
                設定更新の際、毎回空欄となります。ご注意ください。
              </p>
            </div>
            <InputText
              name='password'
              control={control}
              rules={{
                maxLength: {
                  value: 50,
                  message: `50文字以内で入力してください`
                }
              }}
            />
          </div>
          <div className='my-4 flex justify-center'>
            <SubmitButton label={props.labelName} disabled={!canSubmit} />
          </div>
        </form>
      </div>
    </div>
  )
}

type FormLabelProps = {
  label: string
  required?: boolean
  children?: React.ReactNode
}

const FormLabel = ({ label, required = false, children }: FormLabelProps) => {
  const [isModalOpen, setIsModalOpen] = useState(false)
  const toggleModal = (e: any) => {
    if (e.target === e.currentTarget) {
      setIsModalOpen(!isModalOpen)
    }
  }
  const openModal = (e: any) => {
    e.preventDefault()
    setIsModalOpen(true)
  }
  return (
    <label className='block text-sm font-bold'>
      {required && <span className='text-red-500'>*&nbsp;</span>}
      {label}
      {children && (
        <>
          <button onClick={openModal}>
            <QuestionMarkCircleIcon className='base-link ml-1 h-4 w-4' />
          </button>
          {isModalOpen && (
            <Modal close={toggleModal} hideFooter>
              <div className='text-xs'>{children}</div>
            </Modal>
          )}
        </>
      )}
    </label>
  )
}
