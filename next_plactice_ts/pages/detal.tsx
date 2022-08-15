import Link from 'next/link'

export default function Detail() {
  return (
    <div>
      <h1>詳細画面</h1>
      <Link href="/">
          <a>トップ画面へ戻る</a>
      </Link>
    </div>
  )
}