import { Box, Button, Typography } from '@mui/joy';
import TextField from '@mui/material/TextField';
import CameraIcon from '@mui/icons-material/CameraAlt';
import MainButton from '@/components/ui/button/main-button';
import SubButton from '@/components/ui/button/sub-button';
import useUserStore from '@/stores/user-store';
import { useEffect, useState } from 'react';
import CameraModal from '@/components/ui/modal/camera-modal';
import usePostLogoutRequest from '@/features/lobby/api/post-user-logout';

export default function LobbyLayout() {
  const currentUser = useUserStore((state) => state.currentUser);
  const { setTask, setNote, clearUser } = useUserStore();
  const [taskInput, setTaskInput] = useState<string>(currentUser?.task || '');
  const [noteInput, setNoteInput] = useState<string>(currentUser?.note || '');
  const [isCameraOpen, setIsCameraOpen] = useState<boolean>(false);
  const logout = usePostLogoutRequest();

  useEffect(() => {
    if (!currentUser) {
      window.location.href = '/login';
    }
  }, []);

  const handleLogout = async () => {
    await logout.trigger({ iruyanId: currentUser?.iruyanId || '' });
    localStorage.removeItem('user-store');
    clearUser();
    window.location.href = '/login';
  };

  const handleUpdateUser = async () => {
    // if (roomList.data) {
    //   const list = roomList.data.rooms[0].roomId;

    //   if (list) {
    //     setRoomId(list);
    //     console.log('aaa', roomId);
    //   }
    setTask(taskInput);
    setNote(noteInput);
    // const requestData: PostRoomEnterRequest = {
    //   user_id: currentUser.iruyanID,
    //   task: taskInput,
    // };
    try {
      // await roomEntry.entry({ userId: currentUser.iruyanId, roomId: list });
      // if (list) {
      //   window.location.href = `/rooms/:${list}`;
      // }
    } catch (error) {
      console.error('部屋へのエントリーに失敗しました:', error);
    }
  };

  return (
    <Box
      sx={{
        position: 'relative',
        display: 'flex',
        alignItems: 'flex-end',
        justifyContent: 'center',
        height: '100vh',
        backgroundImage: "url('/bg-image/bg_lobby.jpg')",
        backgroundSize: 'cover',
      }}
    >
      <SubButton
        title="ログアウト"
        size="lg"
        onClick={handleLogout}
        sx={{
          position: 'absolute',
          top: '37px',
          right: '200px',
        }}
      />
      <SubButton
        title="来店記録"
        size="lg"
        onClick={() => (window.location.href = '/user')}
        sx={{
          position: 'absolute',
          top: '37px',
          right: '41px',
        }}
      />
      <Box display={'flex'} flexDirection={'column'} alignItems={'center'} mb={20}>
        <Typography level="h4" sx={{ mb: 1, fontSize: '48px' }}>
          ご案内用紙
        </Typography>
        <Typography level="h4" sx={{ mb: 3, color: 'text.secondary' }}>
          {currentUser?.userName}さん、ごゆっくりどうぞ
        </Typography>
        <Box
          sx={{
            display: 'flex',
            justifyContent: 'center',
            gap: '8%',
            width: '100%',
            maxWidth: 840,
          }}
        >
          <Box
            sx={{
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              bgcolor: 'background.default',
              flexDirection: 'column',
              mr: 2,
            }}
          >
            <Typography level="h4">今日のアイコン</Typography>
            <Button
              variant="plain"
              sx={{
                width: '240px',
                height: '240px',
                borderRadius: '50%',
                bgcolor: 'white',
                color: 'text.primary',
                boxShadow: 1,
                '&:hover': {
                  bgcolor: '#f0f0f0',
                },
              }}
              onClick={() => setIsCameraOpen(true)}
            >
              {currentUser?.avatarUrl ? (
                <img
                  src={currentUser.avatarUrl}
                  alt="Avatar"
                  style={{
                    width: '100%',
                    height: '100%',
                    borderRadius: '50%',
                    objectFit: 'cover',
                  }}
                />
              ) : (
                <CameraIcon fontSize="large" />
              )}
            </Button>
          </Box>
          <Box
            sx={{
              display: 'flex',
              alignItems: 'flex-start',
              justifyContent: 'center',
              flexDirection: 'column',
              gap: '44px',
            }}
          >
            <Box
              sx={{
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'flex-start',
                width: '30vw',
              }}
            >
              <Typography level="h4" sx={{ mb: 1 }}>
                ・作業内容
              </Typography>
              <TextField
                placeholder="勉強"
                fullWidth
                value={taskInput}
                onChange={(e) => setTaskInput(e.target.value)}
                sx={{
                  bgcolor: 'white',
                  width: '100%',
                }}
              />
            </Box>
            <Box
              sx={{
                display: 'flex',
                justifyContent: 'center',
                flexDirection: 'column',
                width: '30vw',
              }}
            >
              <Typography level="h4" sx={{ mb: 1 }}>
                ・今日のやる気 / つぶやき
              </Typography>
              <TextField
                placeholder="課題やばい、、よ"
                fullWidth
                value={noteInput}
                onChange={(e) => setNoteInput(e.target.value)}
                sx={{
                  bgcolor: 'white',
                  width: '100%',
                }}
              />
            </Box>
          </Box>
        </Box>
        <MainButton
          title="席に着く"
          type="button"
          maxWidth="240px"
          width="50%"
          component={'div'}
          onClick={handleUpdateUser}
          // disabled={roomList.isLoading}
        />
      </Box>
      <CameraModal isCameraOpen={isCameraOpen} setIsCameraOpen={setIsCameraOpen}></CameraModal>
    </Box>
  );
}
