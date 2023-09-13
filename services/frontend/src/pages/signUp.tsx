import { Link } from 'react-router-dom'
import { useForm, hasLength, matchesField } from '@mantine/form'
import { Center, Container, Paper, Title, Text, Anchor, TextInput, PasswordInput, Button } from '@mantine/core'

const SignUpPage = () => {
  const form = useForm({
    initialValues: {
      username: '',
      password: '',
      confPassword: '',
    },
    validate: {
      username: hasLength({ min: 3, max: 20 }, 'Username must be between 3 and 20 characters'),
      password: hasLength({ min: 8 }, 'Password must have a minimum of 8 characters'),
      confPassword: matchesField('password', 'Confirmation password does not match'),
    }
  })

  return (
    <Center h='95%'>
      <Container>
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
            />
            <PasswordInput 
              label='Password'
              placeholder='Your password'
              mt='md'
              {...form.getInputProps('password')}
            />
            <PasswordInput
              label='Confirm Password'
              placeholder='Confirm password'
              mt='md'
              {...form.getInputProps('confPassword')}
            />
            <Button type='submit' color='teal' mt='xl' fullWidth>Register</Button>
          </form>
        </Paper>
      </Container>
    </Center>
  )
}

export default SignUpPage