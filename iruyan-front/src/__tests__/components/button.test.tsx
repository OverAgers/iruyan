import '@testing-library/jest-dom'
import { render, screen } from '@testing-library/react'
import MainButton from '@/components/ui/button/main-button'

describe('MainButton', () => {
  it('renders button with correct text', () => {
    render(<MainButton title="Click me" type="button" component="button" />)
    expect(screen.getByRole('button', { name: /click me/i })).toBeInTheDocument()
  })

  it('applies custom className', () => {
    render(<MainButton title="Button" type="button" component="button" />)
    const button = screen.getByRole('button')
    expect(button).toHaveClass('MuiButton-root')
  })

  it('handles click events', () => {
    const handleClick = jest.fn()
    render(<MainButton title="Click me" type="button" component="button" onClick={handleClick} />)
    
    const button = screen.getByRole('button')
    button.click()
    
    expect(handleClick).toHaveBeenCalledTimes(1)
  })

  it('can be disabled', () => {
    render(<MainButton title="Disabled Button" type="button" component="button" disabled />)
    const button = screen.getByRole('button')
    expect(button).toBeDisabled()
  })
}) 