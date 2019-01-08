package piplayer

import (
	"fmt"
	"io"
	"os/exec"
)

type PiPlayer struct {
	isRunning      bool
	isPlaying      bool
	hasMediaSource bool
	mediaSource    string
	command        *exec.Cmd
	commandIn      io.WriteCloser
}

func CreatePiPlayer() *PiPlayer {
	return &PiPlayer{
		isRunning:      false,
		isPlaying:      false,
		hasMediaSource: false,
		mediaSource:    "",
		command:        nil,
	}
}

func (p *PiPlayer) waitForPlayerToTerminate() {
	if p.command != nil {
		p.command.Wait()
	}
	fmt.Printf("Player terminated\n")
	p.isPlaying = false
	p.isRunning = false
}

func (p *PiPlayer) launch() {
	p.Quit()
	p.command = exec.Command("omxplayer", p.mediaSource)
	if p.command != nil {
		p.commandIn, _ = p.command.StdinPipe()
		p.command.Start()
		go p.waitForPlayerToTerminate()
		p.isRunning = true
		p.isPlaying = true
	}
}

func (p *PiPlayer) Quit() {
	if p.isRunning {
		p.commandIn.Write([]byte("q"))
	}
}

func (p *PiPlayer) togglePlayPause() {
	if p.isRunning {
		p.commandIn.Write([]byte("p"))
		p.isPlaying = !p.isPlaying
	} else {
		p.launch()
	}
}

func (p *PiPlayer) Play() {
	if !p.isPlaying {
		p.togglePlayPause()
	}
}

func (p *PiPlayer) Pause() {
	if p.isPlaying {
		p.togglePlayPause()
	}
}

func (p *PiPlayer) SetMediaSource(source string) {
	p.mediaSource = source
}
