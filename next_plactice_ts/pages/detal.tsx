import Link from 'next/link'
import useSWR from 'swr'
import { useRouter } from 'next/router'

const fetcher = (url: string) => fetch(url).then(res => res.json())

export default function Detail() {
  const router = useRouter()
  // クエリパラメーターを取得する
  const { id } = router.query
  const { data, error } = useSWR(`http://localhost:8080/detail?id=${id}`, fetcher)
  console.log(data)
  return (
    <div>
      <h1>詳細画面</h1>
      {data ? 
          <div>
            <ul>
              ID
              <li>{data.ID}</li>
              タイトル
              <li>{data.title}</li>
              内容
              <li>{data.content}</li>
              更新日
              <li>{data.UpdatedAt}</li>
              作成日
              <li>{data.CreatedAt}</li>
            </ul>
          </div>
        :
          <p></p>
        }
      <Link href="/">
          <a>トップ画面へ戻る</a>
      </Link>
    </div>
  )
}