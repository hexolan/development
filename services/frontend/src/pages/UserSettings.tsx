import { useOutletContext } from 'react-router-dom'
import { Text } from '@mantine/core'

import { useAppSelector } from '../app/hooks'
import type { UserContext } from '../components/UserLayout'

function UserSettingsPage() {
  const { user } = useOutletContext<UserContext>()
  
  const currentUser = useAppSelector((state) => state.auth.currentUser)
  if (user && (!currentUser || (currentUser.id != user.id && !currentUser.isAdmin))) {
    throw Error('You do not have permission to view that page')
  }

  return (
    <Text>Settings {user.username}</Text>
  )
}

export default UserSettingsPage