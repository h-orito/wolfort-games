import { forwardRef, useImperativeHandle, useRef, useState } from 'react'

type Props = {
  header: string
  children: React.ReactNode
  isOpen?: boolean
  detailsClassName?: string
}

export interface PanelRefHandle {
  open: () => void
}

const Panel = forwardRef<PanelRefHandle, Props>((props: Props, ref: any) => {
  const { header, children, isOpen: initialOpen = true } = props
  const [isOpen, setIsOpen] = useState(initialOpen)

  const detailsRef = useRef<HTMLDetailsElement>(null)

  useImperativeHandle(ref, () => ({
    open() {
      setIsOpen(true)
    }
  }))

  // こうしないとdetails要素のopen属性が変更されない
  const handleClick = (e: any) => {
    e.preventDefault()
    setIsOpen(!isOpen)
  }

  return (
    <>
      <div className='base-border rounded-md border'>
        <details open={isOpen} ref={detailsRef}>
          <summary
            onClick={handleClick}
            className='secondary-background cursor-pointer list-none rounded-t'
          >
            <div className='base-border flex border-b px-3 py-2'>
              <div className='text-md flex-1'>{header}</div>
            </div>
          </summary>
          <div
            className={`primary-text details-content w-full ${
              props.detailsClassName ?? 'p-4'
            }`}
          >
            {children}
          </div>
        </details>
      </div>
    </>
  )
})

export default Panel
