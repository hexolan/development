import { useNavigate } from 'react-router-dom'
import { useForm, hasLength } from '@mantine/form'
import { Center, Container, Paper, Title, TextInput, Textarea, Button } from '@mantine/core'

import { useCreatePanelMutation } from '../app/api/panels'
import type { CreatePanelData } from '../app/types/panels'

const NewPanelPage = () => {
  const navigate = useNavigate()

  const panelForm = useForm<CreatePanelData>({
    initialValues: {
      name: '',
      description: '',
    },
    validate: {
      name: hasLength({ min: 3, max: 20 }, 'Panel name must be between 3 and 20 characters long'),
      description: hasLength({ min: 3, max: 512 }, 'Description must be between 3 and 512 characters'),
    }
  })

  const [createPanel, { isLoading }] = useCreatePanelMutation()
  const submitPanelForm = async (values: CreatePanelData) => {
    await createPanel({
      ...values
    }).unwrap().then((panel) => {
      navigate(`/panel/${panel.name}`)
    }).catch((error) => {
      console.log(error)  // todo: error handling
    })
  }

  return (
    <Center h='95%'>
      <Container>
        <Title align='center' weight={900}>Create a Panel</Title>

        <Paper withBorder shadow='md' radius='md' p={30} mt={30}>
          <form onSubmit={panelForm.onSubmit(submitPanelForm)}>
            <TextInput 
              label='Name'
              placeholder='e.g. music, programming, football'
              {...panelForm.getInputProps('name')}
            />
            <Textarea
              label='Description'
              placeholder='e.g. The place to talk about all things music related...'
              mt={6}
              {...panelForm.getInputProps('description')}
            />

            <Button type='submit' variant='outline' color='teal' mt='xl' disabled={isLoading} fullWidth>Create Panel</Button>
          </form>
        </Paper>
      </Container>
    </Center>
  )
}

export default NewPanelPage