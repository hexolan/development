import { useForm, hasLength } from '@mantine/form'
import { Center, Container, Paper, Title, TextInput, Textarea, Button } from '@mantine/core'

const NewPanelPage = () => {
  const form = useForm({
    initialValues: {
      name: '',
      description: '',
    },
    validate: {
      name: hasLength({ min: 3, max: 20 }, 'Panel name must be between 3 and 20 characters long'),
      description: hasLength({ min: 3, max: 512 }, 'Description must be between 3 and 512 characters'),
    }
  })

  return (
    <Center h='95%'>
      <Container w='25%'>
        <Title align='center' weight={900}>Create a Panel</Title>

        <Paper withBorder shadow='md' radius='md' p={30} mt={30}>
          <form onSubmit={(values) => console.log(values)}>
            <TextInput 
              label='Name'
              placeholder="e.g. music, programming, football"
              {...form.getInputProps('name')}
            />
            <Textarea
              label='Description'
              placeholder="e.g. The place to talk about all things music related..."
              mt={6}
              {...form.getInputProps('description')}
            />
            <Button type='submit' variant='outline' color='teal' mt='xl' fullWidth>Create it!</Button>
          </form>
        </Paper>
      </Container>
    </Center>
  )
}

export default NewPanelPage