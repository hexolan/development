import { Text, Container } from '@mantine/core'

import Feed from '../components/Feed'

const Homepage = () => {
  return (
    <Container mt='xl'>
      <Text>Homepage Text</Text>
      <Feed />
    </Container>
  )
}

export default Homepage