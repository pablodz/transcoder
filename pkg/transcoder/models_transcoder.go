package transcoder

type Transcoder struct {
	MethodT               int                   // the method of transcoding (e.g. 1, 2, 3, etc.)
	MethodR               int                   // the method of resampling (e.g. 1, 2, 3, etc.)
	MethodAdvancedConfigs interface{}           // the specific configuration options for the transcoding method
	SizeBuffer            int                   // the size of the buffer to read from the input file. Default is 1024
	SourceConfigs         TranscoderAudioConfig // the source configuration
	TargetConfigs         TranscoderAudioConfig // the target configuration
	BitDepth              int                   // the bit depth (e.g. 8, 16, 24) Needs to be equal for source and target
	Verbose               bool                  // if true, the transcoding process will be verbose
}

type TranscoderAudioConfig struct {
	Encoding   int // the encoding format (e.g. "UNSIGNED", "SIGNED", "FLOAT", "PCM", "IEEE_FLOAT", "MULAW")
	Endianness int // the endianness of the audio data (e.g. "BIG", "LITTLE")
}
