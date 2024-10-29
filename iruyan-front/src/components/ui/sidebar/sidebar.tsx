export default function Sidebar() {
  return (
    <div className="sidebar">
      <div className="sidebar__header">
        <h2>IRUYAN</h2>
      </div>
      <div className="sidebar__content">
        <ul>
          <li>
            <a href="/room">ルーム</a>
          </li>
          <li>
            <a href="/register">新規登録</a>
          </li>
          <li>
            <a href="/login">ログイン</a>
          </li>
        </ul>
      </div>
    </div>
  );
}