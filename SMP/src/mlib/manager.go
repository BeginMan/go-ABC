package library

import "errors"

//音乐结构体
type MusicEntry struct {
	Id     string //唯一ID
	Name   string //音乐名
	Artist string //艺术家名
	Source string //音乐位置
	Type   string //文件类型(MP3和WAV等)
}

//音乐管理结构体
type MusicManager struct {
	// 使用一个数组切片作为基础存储结构
	// 其他操作其实都只是对这个slice的包装
	musics []MusicEntry
}

//创建音乐实例
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

//如下MusicManager成员方法

//音乐数量
func (m *MusicManager) Len() int {
	return len(m.musics)
}

//通过索引获取对应的音乐
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

//通过名称查找对应的音乐
func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

//添加音乐
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

//通过索引移除对应的音乐
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]

	//从数组切片中删除元素
	if index < len(m.musics)-1 && index > 0 { // 中间元素
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 { //删除仅有的一个元素
		m.musics = make([]MusicEntry, 0)
	} else { // 删除的是最后一个元素
		m.musics = m.musics[:index+1]
	}
	return removedMusic
}

//通过名称移除对应的音乐
func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for i, v := range m.musics {
		if v.Name == name {
			return m.Remove(i)
		}
	}
	return nil
}
