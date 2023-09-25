import { useOutletContext, useNavigate } from 'react-router-dom'
import { useForm, hasLength } from '@mantine/form'
import { Stack, Paper, TextInput, Textarea, Button } from '@mantine/core'

import { useAppSelector } from '../app/hooks'
import { useCreatePanelPostMutation } from '../app/api/posts'
import type { CreatePostData } from '../app/types/posts'
import type { PanelContext } from '../components/PanelLayout'

const NewPanelPostPage = () => {
  const { panel } = useOutletContext<PanelContext>()
  const navigate = useNavigate()

  // Ensure the user is authenticated
  const currentUser = useAppSelector((state) => state.auth.currentUser)
  if (currentUser === null) {
    navigate('/signin')
  }

  const createPostForm = useForm<CreatePostData>({
    initialValues: {
      title: '',
      content: '',
    },
    validate: {
      title: hasLength({ min: 3, max: 512 }, 'Title must be between 3 and 512 characters'),
      content: hasLength({ min: 3, max: 2048 }, 'Content must be between 3 and 2048 characters'),
    }
  })

  const [createPost, { isLoading }] = useCreatePanelPostMutation()
  const submitPost = async (values: CreatePostData) => {
    const postInfo = await createPost({
      panelId: panel.id,
      data: values
    }).unwrap().catch((error) => {
      console.log(error) // todo: error handling & displaying err msg
    })

    // redirect to post page if post was created
    if (postInfo) {
      navigate(`/panel/${panel.name}/posts/${postInfo.id}`)
    }
  }

  return (
    <Paper shadow='md' radius='md' p='lg' withBorder>
      <Stack>
        <form onSubmit={createPostForm.onSubmit(submitPost)}>
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
          
          <Button type='submit' variant='outline' color='teal' mt='xl' disabled={isLoading} fullWidth>
            Create Post
          </Button>
        </form>
      </Stack>
    </Paper>
  )
}

export default NewPanelPostPage