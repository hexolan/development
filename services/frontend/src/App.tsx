import { MantineProvider } from '@mantine/core'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'

import AppLayout from './components/AppLayout'
import IndexPage from './pages/index'
import PanelPage from './pages/panel'
import ErrorPage from './pages/error'

const router = createBrowserRouter([
  {
    element: <AppLayout />,
    errorElement: <AppLayout><ErrorPage /></AppLayout>,
    children: [
      {
        index: true,
        element: <IndexPage />,
      },
      {
        path: '/panel/:panelName',
        element: <PanelPage />,
      }
    ]
  }
])

function App() {
  return (
    <MantineProvider withGlobalStyles withNormalizeCSS>
      <RouterProvider router={router} fallbackElement={<p>Loading...</p>} />
    </MantineProvider>
  );
}

export default App