import Link from 'next/link'
export default function customError() {
  return (
    <div>
      <p>エラーが発生しました。</p>
      <Link href="/">
          <a>トップ画面へ戻る</a>
      </Link>
    </div>
  )
}