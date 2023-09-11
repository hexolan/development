import { apiSlice } from '../api'

export interface CreatePostData {
  title: string;
  content: string;
}

export interface UpdatePostData {
  title?: string;
  content?: string;
}

export const postsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPanelPosts: builder.query({
      query: (name: string) => ({
        url: `/v1/panels/${name}/posts`
      })
    }),

    getPanelPost: builder.query({
      query: (panelName: string, postId: string) => ({
        url: `/v1/panels/${panelName}/posts/${postId}`
      })
    }),

    createPanelPost: builder.mutation({
      query: (panelName: string, data: CreatePostData) => ({
        url: `/v1/panels/${panelName}`,
        method: 'POST',
        body: { ...data }
      })
    }),

    updatePost: builder.mutation({
      query: (id: string, data: UpdatePostData) => ({
        url: `/v1/posts/${id}`,
        method: 'PATCH',
        body: { ...data }
      })
    }),

    deletePost: builder.mutation({
      query: (id: string) => ({
        url: `/v1/posts/${id}`,
        method: 'DELETE'
      })
    })
  })
})

export const { useGetPanelPostsQuery, useGetPanelPostQuery, useCreatePanelPostMutation, useUpdatePostMutation, useDeletePostMutation } = postsApiSlice