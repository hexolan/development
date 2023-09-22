import { useOutletContext } from 'react-router-dom'
import { Text } from '@mantine/core'

import type { UserContext } from '../components/UserLayout'

function UserAboutPage() {
  const { user } = useOutletContext<UserContext>()

  return (
    <Text>About {user.username}</Text>
  )
}

export default UserAboutPage