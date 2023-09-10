import { Text } from '@mantine/core'
import { useParams } from 'react-router-dom'

type UserPageParams = {
  username: string;
}

function UserPage() {
  const { username } = useParams<UserPageParams>();

  return (
    <Text>User - {username}</Text>
  )
}

export default UserPage