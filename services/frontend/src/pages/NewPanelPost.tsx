import { useForm, hasLength } from '@mantine/form'
import { Stack, Paper, Text, TextInput, Textarea, Button } from '@mantine/core'

const NewPanelPostPage = () => {
  const createPostForm = useForm({
    initialValues: {
      title: '',
      content: '',
    },
    validate: {
      title: hasLength({ min: 3, max: 20 }, 'Title must be between X and Y characters'),  // todo: change min and max
      content: hasLength({ min: 3, max: 512 }, 'Content must be between X and Y characters'),  // todo: change min and max
    }
  })

  return (
    <Paper shadow='md' radius='md' p='lg' withBorder>
      <Stack>
        <Text size='xl' weight={500} align='center'>New Post</Text>

        <form onSubmit={(values) => console.log(values)}>
          <TextInput 
            label='Title'
            placeholder='Post Title'
            {...createPostForm.getInputProps('title')}
          />
          <Textarea
            label='Content'
            placeholder='Post Content'
            mt={6}
            {...createPostForm.getInputProps('content')}
          />
          <Button type='submit' variant='outline' color='teal' mt='xl' fullWidth>Create Post</Button>
        </form>
      </Stack>
    </Paper>
  )
}

export default NewPanelPostPage