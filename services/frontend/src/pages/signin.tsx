import { Link } from 'react-router-dom'
import { Center, Container, Paper, Title, Text, Anchor, TextInput, PasswordInput, Button } from '@mantine/core'

const SignInPage = () => {
  return (
    <Center h='95%'>
      <Container w='25%'>
        <Title align='center' weight={900}>Sign In</Title>
        <Text color='dimmed' size='sm' align='center' mt={5}>
          Do not have an account yet?{' '}
          <Anchor size='sm' component={Link} to='/signup'>Sign up</Anchor> instead.
        </Text>

        <Paper withBorder shadow='md' radius='md' p={30} mt={30}>
          <TextInput label='Username' placeholder="hexolan" required />
          <PasswordInput label='Password' placeholder='Your password' mt='md' required />
          <Button color='teal' fullWidth mt='xl'>
            Sign In
          </Button>
        </Paper>
      </Container>
    </Center>
  )
}

export default SignInPage