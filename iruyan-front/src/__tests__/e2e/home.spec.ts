import { test, expect } from '@playwright/test'

test.describe('Home Page', () => {
  test('should load the home page', async ({ page }) => {
    await page.goto('/')
    
    // Check if the page loads without errors
    await expect(page).toHaveTitle(/iruyan/)
  })

  test('should have proper navigation', async ({ page }) => {
    await page.goto('/login')
    
    // Check if navigation elements are present
    // You can add more specific checks based on your actual navigation structure
    // await expect(page.locator('nav')).toBeVisible()
  })

  test('should be responsive', async ({ page }) => {
    await page.goto('/')
    
    // Test mobile viewport
    await page.setViewportSize({ width: 375, height: 667 })
    await expect(page.locator('body')).toBeVisible()
    
    // Test desktop viewport
    await page.setViewportSize({ width: 1920, height: 1080 })
    await expect(page.locator('body')).toBeVisible()
  })
}) 