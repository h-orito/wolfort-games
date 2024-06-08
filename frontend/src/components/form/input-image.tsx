import { Dispatch, SetStateAction, useRef, useState } from 'react'
import DangerButton from '../button/danger-button'
import PrimaryButton from '../button/primary-button'

interface Props {
  label?: string
  name: string
  setImages: Dispatch<SetStateAction<File[]>>
  defaultImageUrl?: string | null
  maxFileKByte?: number
  disabled?: boolean
}

const allowImageTypes = [
  'image/jpeg',
  'image/png',
  'image/jpg',
  'image/gif',
  'image/webp'
]

export default function InputImage({
  label,
  name,
  setImages,
  defaultImageUrl,
  maxFileKByte = 1024,
  disabled = false
}: Props) {
  const [errorMessage, setErrorMessage] = useState('')
  const [previewImageUrl, setPreviewImageUrl] = useState(defaultImageUrl)
  const inputRef = useRef<HTMLInputElement>(null!)

  const onProfileButtonClick = (e: any) => {
    e.preventDefault()
    inputRef.current.click()
  }

  const handleFile = async (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files === null || event.target.files.length === 0) {
      return
    }
    const file = event.target.files[0]
    // 初期化することで同じファイルを連続で選択してもonChagngeが発動するように設定し、画像をキャンセルしてすぐに同じ画像を選ぶ動作に対応
    event.target.value = ''
    setErrorMessage('')

    if (!allowImageTypes.includes(file.type)) {
      setErrorMessage('jpeg, jpg, png, gif, webpのファイルを選択してください')
      return
    }

    if (file.size > maxFileKByte * 1024) {
      setErrorMessage(`${maxFileKByte}kByte以下のファイルを選択してください`)
      return
    }

    setImages([file])
    setPreviewImageUrl(URL.createObjectURL(file))
  }

  const handleCancel = (e: any) => {
    e.preventDefault()
    setErrorMessage('')
    setImages([])
    setPreviewImageUrl('')
  }

  return (
    <div>
      {label && <label className='block text-xs font-bold'>{label}</label>}
      {previewImageUrl && (
        <img
          className='mb-2 w-32'
          src={previewImageUrl}
          alt={label ? label : '画像'}
        />
      )}

      <input
        type='file'
        name={name}
        id={name}
        ref={inputRef}
        accept='image/*'
        onChange={handleFile}
        hidden
      />
      <PrimaryButton click={onProfileButtonClick} disabled={disabled}>
        画像を選択
      </PrimaryButton>
      {previewImageUrl != null && (
        <DangerButton className='ml-2' click={(e: any) => handleCancel(e)}>
          削除
        </DangerButton>
      )}
      {errorMessage && <p className='text-xs text-red-500'>{errorMessage}</p>}
    </div>
  )
}
