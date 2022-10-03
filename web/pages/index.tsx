import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import CenteredLayout from '../components/CenteredLayout'
import useSession from '../hooks/useSession'

const Home: NextPage = () => {
  const router = useRouter()
  const [session, setSession, sessionLoading, sessionError] = useSession()

  if (sessionLoading)
    return <div>Spinner</div>
  if (sessionError)
    return <div>{String(sessionError)}</div>

  if (!session.loggedIn)
    return (
        <CenteredLayout title="EpyTodo Remix">
          <button className="text-dark" type="button" onClick={() => router.push('/login')}>
            connecte toi grobatar
          </button>
        </CenteredLayout>
    )

  return (
    <div>
      <div>Hello {session.user.username}</div>
    </div>
  )
}

export default Home
