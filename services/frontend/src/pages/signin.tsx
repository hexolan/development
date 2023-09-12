import { Link } from 'react-router-dom'
import { useForm, hasLength } from '@mantine/form'
import { Center, Container, Paper, Title, Text, Anchor, TextInput, PasswordInput, Button } from '@mantine/core'

const SignInPage = () => {
  const form = useForm({
    initialValues: {
      username: '',
      password: '',
    },
    validate: {
      username: hasLength({ min: 3, max: 20 }, 'Invalid username'),
      password: hasLength({ min: 8 }, 'Invalid password'),
    }
  })

  return (
    <Center h='95%'>
      <Container w='25%'>
        <Title align='center' weight={900}>Sign In</Title>
        <Text color='dimmed' size='sm' align='center' mt={5}>
          Do not have an account yet?{' '}
          <Anchor size='sm' component={Link} to='/signup'>Sign up</Anchor> instead.
        </Text>

        <Paper withBorder shadow='md' radius='md' p={30} mt={30}>
          <form onSubmit={form.onSubmit((values) => console.log(values))}>
            <TextInput 
              label='Username'
              placeholder="hexolan" 
              {...form.getInputProps('username')}
            />
            <PasswordInput 
              label='Password' 
              placeholder='Your password' 
              mt='md'
              {...form.getInputProps('password')}
            />
            <Button type='submit' color='teal' mt='xl' fullWidth>Sign In</Button>
          </form>
        </Paper>
      </Container>
    </Center>
  )
}

export default SignInPage