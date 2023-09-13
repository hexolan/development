import { apiSlice } from '../api'

// import type { Post } from '../types'

interface GetPanelPostRequest {
  panelName: string;
  postId: string;
}

interface CreatePanelPostRequest {
  panelName: string;
  data: CreatePostData;
}

export interface CreatePostData {
  title: string;
  content: string;
}

interface UpdatePostRequest {
  postId: string;
  data: UpdatePostData;
}

export interface UpdatePostData {
  title?: string;
  content?: string;
}

export const postsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPanelPosts: builder.query({
      query: (panelName: string) => `/v1/panels/${panelName}/posts`,
      // transformResponse: (response: Object[]) => { return response.data }, // todo (convert times to ISO strings)
    }),

    getPanelPost: builder.query({
      query: (req: GetPanelPostRequest) => ({ url: `/v1/panels/${req.panelName}/posts/${req.postId}` }),
      // transformResponse: (response: Post) => { return response.data },  // todo: return as post (same as above)
    }),

    createPanelPost: builder.mutation({
      query: (req: CreatePanelPostRequest) => ({
        url: `/v1/panels/${req.panelName}`,
        method: 'POST',
        body: { ...req.data },
      }),
      // transformResponse: (response: Post) => { return response.data },  // todo: return as post (same as above)
    }),

    updatePost: builder.mutation({
      query: (req: UpdatePostRequest) => ({
        url: `/v1/posts/${req.postId}`,
        method: 'PATCH',
        body: { ...req.data },
      }),
      // transformResponse: (response: Post) => { return response.data },  // todo: return as post (same as above)
    }),

    deletePost: builder.mutation({
      query: (postId: string) => ({
        url: `/v1/posts/${postId}`,
        method: 'DELETE',
      }),
    }),
  })
})

export const { useGetPanelPostsQuery, useGetPanelPostQuery, useCreatePanelPostMutation, useUpdatePostMutation, useDeletePostMutation } = postsApiSlice