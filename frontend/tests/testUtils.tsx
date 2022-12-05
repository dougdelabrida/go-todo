import { ChakraProvider } from '@chakra-ui/react'
import { render } from '@testing-library/react'

export const renderWithChakra = (ui: any) => {
  const Wrapper = ({ children }: any) => <ChakraProvider>{children}</ChakraProvider>
  return render(ui, { wrapper: Wrapper })
}
