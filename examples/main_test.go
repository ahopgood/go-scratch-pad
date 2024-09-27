package examples_test

import (
	"bytes"
	"context"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"log"
	"log/slog"
	"strings"
)

var _ = Describe("Main", func() {
	When("log", func() {
		It("should log dates", func() {
			var buf bytes.Buffer
			log.SetOutput(&buf)

			//2024/09/22 14:16:52 DEBUG test message
			slog.Info("test message")
			actualLogLevel := strings.Fields(buf.String())[0]

			Expect(actualLogLevel).To(MatchRegexp("[[:digit:]]{4}/[[:digit:]]{2}/[[:digit:]]{2}"))
		})

		It("should log timestamps in microseconds", func() {
			var buf bytes.Buffer
			log.SetOutput(&buf)
			log.SetFlags(log.LstdFlags | log.Lmicroseconds)

			//2024/09/24 20:14:43.812102 INFO test message
			slog.Info("test message")
			//fmt.Println(buf.String())
			actualLogLevel := strings.Fields(buf.String())[1]
			Expect(actualLogLevel).To(MatchRegexp("^[[:digit:]]{2}:[[:digit:]]{2}:[[:digit:]]{2}\\.[[:digit:]]{6}$"))
		})

		It("should log short style code locations", func() {
			var buf bytes.Buffer
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.SetOutput(&buf)

			//2024/09/22 14:16:52 DEBUG test message
			slog.Info("test message")
			fmt.Println(buf.String())
			actualLogLevel := strings.Fields(buf.String())[2]
			Expect(actualLogLevel).To(MatchRegexp("[[:alpha:]]+_test\\.go:[[:digit:]]+:"))
		})

		It("should log long style code locations", func() {
			var buf bytes.Buffer
			log.SetFlags(log.LstdFlags | log.Llongfile)
			log.SetOutput(&buf)

			//2024/09/22 14:16:52 DEBUG test message
			slog.Info("test message")
			fmt.Println(buf.String())
			actualLogLevel := strings.Fields(buf.String())[2]
			Expect(actualLogLevel).To(MatchRegexp("examples/[[:alpha:]]+_test\\.go:[[:digit:]]+:"))
		})

		It("should log with key-value pair", func() {
			var buf bytes.Buffer
			log.SetFlags(log.LstdFlags)
			log.SetOutput(&buf)

			//2024/09/22 14:16:52 DEBUG test message
			slog.Info("test message", "mykey", "myvalue")

			fmt.Println(buf.String())
			actualLogLevel := strings.Fields(buf.String())[5]
			Expect(actualLogLevel).To(Equal("mykey=myvalue"))
		})

		It("should not be in json", func() {
			var buf bytes.Buffer
			log.SetFlags(log.LstdFlags)
			log.SetOutput(&buf)

			//2024/09/22 14:16:52 DEBUG test message
			slog.Info("test message")
			actualLogLevel := strings.Fields(buf.String())[3:5]
			Expect(actualLogLevel).To(ContainElements("test", "message"))
		})

		DescribeTable("should default to Info level",
			func(level slog.Level, expectedLevel string, shouldLog bool) {
				var buf bytes.Buffer
				log.SetFlags(log.LstdFlags)
				log.SetOutput(&buf)

				slog.Log(context.Background(), level, "test message")
				//outputs:
				//2024/09/22 14:16:52 DEBUG test message
				//fmt.Println(buf.String())

				if shouldLog {
					actualLogLevel := strings.Fields(buf.String())[2]
					Expect(actualLogLevel).To(ContainSubstring(expectedLevel))
				} else {
					Expect(buf.String()).To(BeEmpty())
				}

			},
			Entry(
				fmt.Sprintf("Log level set to %d", slog.LevelDebug), slog.LevelDebug, "DEBUG", false),
			Entry(
				fmt.Sprintf("Log level set to %d", slog.LevelInfo), slog.LevelInfo, "INFO", true),
			Entry(
				fmt.Sprintf("Log level set to %d", slog.LevelWarn), slog.LevelWarn, "WARN", true),
			Entry(
				fmt.Sprintf("Log level set to %d", slog.LevelError), slog.LevelError, "ERROR", true),
		)
	})
})
