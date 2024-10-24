"use client"
export default function RegisterPage() {
  return (
    <div>
      <h1>Register</h1>
      <form>
        <input type="text" placeholder="Username" />
        <input type="password" placeholder="Password" />
        <button type="submit">Register</button>
      </form>
    </div>
  );
}