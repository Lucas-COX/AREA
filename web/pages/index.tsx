import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import CenteredLayout from '../components/CenteredLayout'
import useSession from '../hooks/useSession'

const Home: NextPage = () => {
  const router = useRouter()
  const [session, loading, error] = useSession()

  if (loading)
    return <div>Spinner</div>

  if (error)
    return <div>{String(error)}</div>

  if (session.user == null)
    router.push("/login")

  return (
    <CenteredLayout title={"AREA"}>
      {loading ?
        <div>Spinner</div> :
        <div>Hello {session?.user?.username}</div>
      }
    </CenteredLayout>
  )
}

export default Home
