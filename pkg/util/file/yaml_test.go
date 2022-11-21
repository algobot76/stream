package file_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/devstream-io/devstream/pkg/util/file"
)

var _ = Describe("YamlSequenceNode struct", func() {
	var n *file.YamlSequenceNode
	Context("IsEmpty method", func() {
		When("StrOrigin is not empty", func() {
			BeforeEach(func() {
				n = &file.YamlSequenceNode{
					StrOrigin: "test",
				}
			})
			It("should return false", func() {
				Expect(n.IsEmpty()).Should(BeFalse())
			})
		})
		When("StrArray is not empty", func() {
			BeforeEach(func() {
				n = &file.YamlSequenceNode{
					StrArray: []string{"test"},
				}
			})
			It("should return false", func() {
				Expect(n.IsEmpty()).Should(BeFalse())
			})
		})
		When("StrArray is empty", func() {
			BeforeEach(func() {
				n = &file.YamlSequenceNode{}
			})
			It("should return true", func() {
				Expect(n.IsEmpty()).Should(BeTrue())
			})
		})
	})
})

var _ = Describe("GetYamlNodeArrayByPath", func() {
	var (
		dst, src *file.YamlSequenceNode
	)
	When("src is nil", func() {
		BeforeEach(func() {
			dst = &file.YamlSequenceNode{
				StrOrigin: "test_dst",
			}
			src = nil
		})
		It("should return dst content", func() {
			Expect(file.MergeYamlNode(dst, src)).Should(Equal(dst))
		})
	})
	When("dst is nil", func() {
		BeforeEach(func() {
			src = &file.YamlSequenceNode{
				StrOrigin: "test_src",
			}
			dst = nil
		})
		It("should return src content", func() {
			Expect(file.MergeYamlNode(dst, src)).Should(Equal(src))
		})
	})
	When("dst and src have contents", func() {
		BeforeEach(func() {
			src = &file.YamlSequenceNode{
				StrOrigin: "test_src",
				StrArray:  []string{"test_src_array"},
			}
			dst = &file.YamlSequenceNode{
				StrOrigin: "test_dst",
				StrArray:  []string{"test_dst_array"},
			}
		})
		It("should merge content", func() {
			result := file.MergeYamlNode(dst, src)
			Expect(result.StrOrigin).Should(Equal("test_dst\ntest_src"))
			Expect(result.StrArray).Should(Equal([]string{
				"test_dst_array", "test_src_array",
			}))
		})
	})
})

var _ = Describe("GetYamlNodeArrayByPath", func() {
	var (
		yamlPath string
		testData []byte
	)
	BeforeEach(func() {
		testData = []byte(`
tests:
  - name: plugin1
    instanceID: default
    options:
      key1: [[ var1 ]]
  - name: plugin2
    instanceID: ins2
    options:
      key1: value1
      key2: [[ var2 ]]`)
	})
	When("yaml path is not valid", func() {
		BeforeEach(func() {
			yamlPath = "not_valid_path"
		})
		It("should return error", func() {
			_, err := file.GetYamlNodeArrayByPath(testData, yamlPath)
			Expect(err).Error().Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("invalid path string"))
		})
	})
	When("yaml path not exist", func() {
		BeforeEach(func() {
			yamlPath = "$.field"
		})
		It("should return nil", func() {
			node, err := file.GetYamlNodeArrayByPath(testData, yamlPath)
			Expect(err).Error().ShouldNot(HaveOccurred())
			Expect(node).Should(BeNil())
		})
	})
	When("node is valid sequenceNode", func() {
		var expectStr string
		BeforeEach(func() {
			yamlPath = "$.tests[*]"
			expectStr = `  - name: plugin1
    instanceID: default
    options:
      key1: [[var1]]
  - name: plugin2
    instanceID: ins2
    options:
      key1: value1
      key2: [[var2]]`
		})
		It("should return sequenceNode", func() {
			node, err := file.GetYamlNodeArrayByPath(testData, yamlPath)
			Expect(err).Error().ShouldNot(HaveOccurred())
			Expect(node).ShouldNot(BeNil())
			Expect(expectStr).Should(Equal(node.StrOrigin))
			nodeArray := node.StrArray
			Expect(len(nodeArray)).Should(Equal(2))
			Expect(nodeArray[0]).Should(Equal("    name: plugin1\n    instanceID: default\n    options:\n      key1: [[var1]]"))
			Expect(nodeArray[1]).Should(Equal("    name: plugin2\n    instanceID: ins2\n    options:\n      key1: value1\n      key2: [[var2]]"))
		})
	})
	When("yaml data array is not valid", func() {
		BeforeEach(func() {
			testData = []byte(`
tests:
  - name: plugin1
    instanceID: default
    options:
      key1: ggg`)
		})
		It("should return error", func() {
			_, err := file.GetYamlNodeArrayByPath(testData, yamlPath)
			Expect(err).Error().ShouldNot(HaveOccurred())
		})
	})
	When("node is not sequence", func() {
		BeforeEach(func() {
			testData = []byte(`tests:
  name: plugin1
  instanceID: default
  options:
    key1: [[ var1 ]]`)
			yamlPath = "$.tests.name"
		})
		It("should return error", func() {
			_, err := file.GetYamlNodeArrayByPath(testData, yamlPath)
			Expect(err).Error().Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("is not valid sequenceNode"))
		})
	})
})

var _ = Describe("GetYamlNodeStrByPath", func() {
	var (
		testData []byte
		yamlPath string
	)
	BeforeEach(func() {
		yamlPath = "$.tests"
		testData = []byte(`
tests:
  name: plugin1
  instanceID: default
  options:
    key1: [[ var ]]`)
	})
	It("should return error", func() {
		node, err := file.GetYamlNodeStrByPath(testData, yamlPath)
		Expect(err).Error().ShouldNot(HaveOccurred())
		Expect(node).Should(Equal("  name: plugin1\n  instanceID: default\n  options:\n    key1: [[var]]"))
	})
})
