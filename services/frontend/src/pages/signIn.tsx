import { Link, useNavigate } from 'react-router-dom'
import { useForm, hasLength } from '@mantine/form'
import { Center, Container, Paper, Title, Text, Anchor, TextInput, PasswordInput, Button } from '@mantine/core'

import { useAppDispatch } from '../app/hooks'
import { setSignedIn } from '../app/features/auth'
import { useSignInMutation } from '../app/features/authApi'

interface SignInFormValues {
  username: string;
  password: string;
}

function SignInPage() {
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

  // const dispatch = useAppDispatch()
  // const [attemptSignIn, { isLoading }] = useSignInMutation()
  
  const dispatch = useAppDispatch()
  const [requestSignIn] = useSignInMutation()

  const formSignIn = async (values: SignInFormValues) => {
    // TODO: also check that the user is not already signed in
    // (for ability to view this form - otherwise msg and attempted redirect to homepage)
    let authInfo = await requestSignIn(values).unwrap()
      .catch((error: Error) => console.error('failed', error) )
    dispatch(setSignedIn(authInfo))

    // todo: loading spinner
    // error handling
    // passing token to store in state
    // etc...
    /*
    await attemptSignIn(values)
      .unwrap()
      .then((payload) => console.log('success', payload))
      .catch((error) => console.error('failed', error))

    dispatch(
      setSignedIn({
        token: 'abc',
        currentUser: {
          id: 'test',
          username: values.username,
          createdAt: new Date().toISOString()
        } 
      })
    )
    */

    // attempt to sign in
    
    // set state to pending (render spinner)

    // succesful authentication -> redirection
    
    console.log(values)
    navigate('/')
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
              placeholder="hexolan" 
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