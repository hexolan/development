import { useOutletContext } from 'react-router-dom'
import { Text } from '@mantine/core'

import type { UserContext } from '../components/UserLayout'

function UserSettingsPage() {
  const { user } = useOutletContext<UserContext>()

  return (
    <Text>Settings {user.username}</Text>
  )
}

export default UserSettingsPage