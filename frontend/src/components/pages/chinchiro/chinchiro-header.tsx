import Link from 'next/link'

const ChinchiroHeader = () => {
  return (
    <div className='w-screen bg-slate-200 p-4 text-left'>
      <Link href='/chinchiro'>
        <h1>チンチロ</h1>
      </Link>
    </div>
  )
}
export default ChinchiroHeader
