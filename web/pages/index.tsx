import type { GetServerSidePropsContext } from 'next'
import AppLayout from '../components/AppLayout'
import { getSession } from '../lib/session'
import { withSession } from '../config/withs'

export default function Home({ session }: HomeProps) {
  return (
    <AppLayout type="centered">
        <div>Hello {session.user?.username}</div>
    </AppLayout>
  )
}

export interface HomeProps extends withSession {}

export async function getServerSideProps(context: GetServerSidePropsContext) {
  const session = await getSession(context)
  if (session.authenticated == false) {
    return {
      redirect: {
        destination: '/login',
        permanent: false,
      },
    }
  }
  return {
    props: { session }
  }
}
