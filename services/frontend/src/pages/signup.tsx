import { Link } from 'react-router-dom'
import { Center, Container, Paper, Title, Text, Anchor, TextInput, PasswordInput, Button } from '@mantine/core'

const SignUpPage = () => {
  return (
    <Center h='95%'>
      <Container w='25%'>
        <Title align='center' weight={900}>Sign Up</Title>
        <Text color='dimmed' size='sm' align='center' mt={5}>
          Already have an account?{' '}
          <Anchor size='sm' component={Link} to='/signin'>Sign in</Anchor> instead.
        </Text>

        <Paper withBorder shadow='md' radius='md' p={30} mt={30}>
          <TextInput label='Username' placeholder="hexolan" required />
          <PasswordInput label='Password' placeholder='Your password' mt='md' required />
          <PasswordInput label='Confirm Password' placeholder='Confirm password' mt='md' required />
          <Button color='teal' fullWidth mt='xl'>Register</Button>
        </Paper>
      </Container>
    </Center>
  )
}

export default SignUpPage