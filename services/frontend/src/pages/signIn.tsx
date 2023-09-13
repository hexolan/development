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
    
    // todo: status on auth state (e.g. idle, pending, authed)
    // from this... displaying a spinner when pending

    // Attempt to authenticate the user.
    let authInfo = await attemptLogin(values).unwrap().catch(
      (error) => {
        // todo: proper error handling
        // errors with no data returned (e.g. API offline - go to Uh oh page)
        setErrorMsg(error.data.msg)
      }
    )

    // Check if authentication was succesful.
    if (authInfo) {
      // Redirect to homepage.
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
              my='md'
              {...signinForm.getInputProps('password')}
            />

            {errorMsg !== '' ? <Text color='red' align='center' mb='md'>{'Error: ' + errorMsg}</Text> : null}

            <Button type='submit' color='teal' fullWidth>Sign In</Button>
          </form>
        </Paper>
      </Container>
    </Center>
  )
}

export default SignInPage