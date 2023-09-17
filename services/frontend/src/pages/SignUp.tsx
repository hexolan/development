import { useState } from 'react'

import { Link, useNavigate } from 'react-router-dom'
import { useForm, hasLength, matchesField } from '@mantine/form'
import { Center, Container, Paper, Title, Text, Anchor, TextInput, PasswordInput, Button } from '@mantine/core'

import { useRegisterUserMutation } from '../app/api/users'

type RegistrationFormValues = {
  username: string;
  password: string;
  confPassword: string;
}

const SignUpPage = () => {
  const [errorMsg, setErrorMsg] = useState('')
  const navigate = useNavigate()

  const registrationForm = useForm<RegistrationFormValues>({
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

  const [registerUser] = useRegisterUserMutation()
  const submitRegistrationForm = async (values: RegistrationFormValues) => {
    const req = {username: values.username, password: values.password}
    const authInfo = await registerUser(req).unwrap().catch(
      (error) => {
        // todo: proper error handling
        // errors with no data returned (e.g. API offline - go to Uh oh page)
        if (!error.data) {
          console.log(error)
          setErrorMsg('unable to access api')
        } else {
          setErrorMsg(error.data.msg)
        }
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
        <Title align='center' weight={900}>Sign Up</Title>
        <Text color='dimmed' size='sm' align='center' mt={5}>
          Already have an account?{' '}
          <Anchor size='sm' component={Link} to='/signin'>Sign in</Anchor> instead.
        </Text>

        <Paper withBorder shadow='md' radius='md' p={30} mt={30}>
          <form onSubmit={registrationForm.onSubmit(submitRegistrationForm)}>
            <TextInput 
              label='Username'
              placeholder='Your username'
              {...registrationForm.getInputProps('username')}
            />
            <PasswordInput 
              label='Password'
              placeholder='Your password'
              my='md'
              {...registrationForm.getInputProps('password')}
            />
            <PasswordInput
              label='Confirm Password'
              placeholder='Confirm password'
              my='md'
              {...registrationForm.getInputProps('confPassword')}
            />

            {errorMsg !== '' && <Text color='red' align='center' mb='md'>{'Error: ' + errorMsg}</Text>}

            <Button type='submit' color='teal' fullWidth>Register</Button>
          </form>
        </Paper>
      </Container>
    </Center>
  )
}

export default SignUpPage