import { useState } from 'react'

import { Link, useNavigate } from 'react-router-dom'
import { useForm, hasLength } from '@mantine/form'
import { Center, Container, Paper, Title, Text, Anchor, TextInput, PasswordInput, Button } from '@mantine/core'

import { useLoginMutation } from '../app/api/auth'

interface SignInFormValues {
  username: string;
  password: string;
}

function SignInPage() {
  const [errorMsg, setErrorMsg] = useState('')
  const navigate = useNavigate()

  const signinForm = useForm<SignInFormValues>({
    initialValues: {
      username: '',
      password: '',
    },
    validate: {
      username: hasLength({ min: 3, max: 20 }, 'Invalid username'),
      password: hasLength({ min: 3 }, 'Invalid password'),
    }
  })

  const [attemptLogin, { isLoading }] = useLoginMutation()
  const formSignIn = async (values: SignInFormValues) => {
    // TODO: also check that the user is not already signed in
    // status on auth state (e.g. idle, pending, authed)

    // attempt to sign in
    console.log(isLoading)
    let authInfo = await attemptLogin(values).unwrap()
    .catch(
      (error) => {
        console.error('failed', error)
        setErrorMsg('Error: ' + error.data.msg)
      }
    )
    console.log(isLoading)

    // succesful authentication -> redirection
    if (authInfo) {
      navigate('/')
    }
  }

  return (
    <Center h='95%'>
      <Container>
        <Title align='center' weight={900}>Sign In</Title>
        <Text color='dimmed' size='sm' align='center' mt={5}>
          Do not have an account yet?{' '}
          <Anchor size='sm' component={Link} to='/signup'>Sign up</Anchor> instead.
        </Text>

        <Text>{errorMsg}</Text>

        <Paper withBorder shadow='md' radius='md' p={30} mt={30}>
          <form onSubmit={signinForm.onSubmit(formSignIn)}>
            <TextInput 
              label='Username'
              placeholder="Your username" 
              {...signinForm.getInputProps('username')}
            />
            <PasswordInput 
              label='Password' 
              placeholder='Your password' 
              mt='md'
              {...signinForm.getInputProps('password')}
            />
            <Button type='submit' color='teal' mt='xl' fullWidth>Sign In</Button>
          </form>
        </Paper>
      </Container>
    </Center>
  )
}

export default SignInPage