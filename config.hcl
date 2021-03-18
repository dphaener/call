command "ls" {
  use = "ls"
  short = "its ls dummy"
  expanded = "ls"
  args = []
}

command "tl" {
  use = "tl"
  short = "List tmux sessions"
  expanded = "tmux ls"
  args = []
}
