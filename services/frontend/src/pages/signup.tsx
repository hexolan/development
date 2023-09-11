import { Link } from 'react-router-dom'
import { useForm, matchesField } from '@mantine/form'
import { Center, Container, Paper, Title, Text, Anchor, TextInput, PasswordInput, Button } from '@mantine/core'

const SignUpPage = () => {
  const form = useForm({
    initialValues: {
      username: '',
      password: '',
      confPassword: '',
    },
    validate: {
      confPassword: matchesField('password', 'The passwords do not match'),
    },
  })

  return (
    <Center h='95%'>
      <Container w='25%'>
        <Title align='center' weight={900}>Sign Up</Title>
        <Text color='dimmed' size='sm' align='center' mt={5}>
          Already have an account?{' '}
          <Anchor size='sm' component={Link} to='/signin'>Sign in</Anchor> instead.
        </Text>

        <Paper withBorder shadow='md' radius='md' p={30} mt={30}>
          <form onSubmit={form.onSubmit((values) => console.log(values))}>
            <TextInput 
              label='Username'
              placeholder="hexolan"
              {...form.getInputProps('username')}
              required
            />
            <PasswordInput 
              label='Password'
              placeholder='Your password'
              mt='md'
              {...form.getInputProps('password')}
              required
            />
            <PasswordInput
              label='Confirm Password'
              placeholder='Confirm password'
              mt='md'
              {...form.getInputProps('confPassword')}
              required
            />
            <Button type='submit' color='teal' mt='xl' fullWidth>Register</Button>
          </form>
        </Paper>
      </Container>
    </Center>
  )
}

export default SignUpPage