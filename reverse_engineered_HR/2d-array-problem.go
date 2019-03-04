package hrreversed

import "time"

type HRReversed struct {
	Community struct {
		GlobalMessage string `json:"globalMessage"`
		Domains       struct {
			List []struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Slug     string `json:"slug"`
				Chapters []struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Slug   string `json:"slug"`
					Hidden bool   `json:"hidden"`
				} `json:"chapters"`
			} `json:"list"`
			Dict struct {
				Tutorials struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						Three0DaysOfCode struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"30-days-of-code"`
						One0DaysOfStatistics struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"10-days-of-statistics"`
						One0DaysOfJavascript struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"10-days-of-javascript"`
					} `json:"chapterDict"`
				} `json:"tutorials"`
				Algorithms struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						Warmup struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"warmup"`
						Implementation struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"implementation"`
						Strings struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"strings"`
						ArraysAndSorting struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"arrays-and-sorting"`
						Search struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"search"`
						GraphTheory struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"graph-theory"`
						Greedy struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"greedy"`
						DynamicProgramming struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"dynamic-programming"`
						ConstructiveAlgorithms struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"constructive-algorithms"`
						BitManipulation struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"bit-manipulation"`
						Recursion struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"recursion"`
						GameTheory struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"game-theory"`
						NpCompleteProblems struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"np-complete-problems"`
						AlgoDebugging struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"algo-debugging"`
					} `json:"chapterDict"`
				} `json:"algorithms"`
				DataStructures struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						Arrays struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"arrays"`
						LinkedLists struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"linked-lists"`
						Trees struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"trees"`
						BalancedTrees struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"balanced-trees"`
						Stacks struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"stacks"`
						Queues struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"queues"`
						Heap struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"heap"`
						DisjointSet struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"disjoint-set"`
						MultipleChoice struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"multiple-choice"`
						Trie struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"trie"`
						DataStructures struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"data-structures"`
					} `json:"chapterDict"`
				} `json:"data-structures"`
				Mathematics struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						Fundamentals struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"fundamentals"`
						NumberTheory struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"number-theory"`
						Combinatorics struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"combinatorics"`
						Algebra struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"algebra"`
						Geometry struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"geometry"`
						Probability struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"probability"`
						LinearAlgebraFoundations struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"linear-algebra-foundations"`
					} `json:"chapterDict"`
				} `json:"mathematics"`
				C struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						CIntroduction struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"c-introduction"`
						CConditionalsAndLoops struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"c-conditionals-and-loops"`
						CArraysAndStrings struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"c-arrays-and-strings"`
						CFunctions struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"c-functions"`
						CStructsAndEnums struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"c-structs-and-enums"`
					} `json:"chapterDict"`
				} `json:"c"`
				Ai struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						AiIntroduction struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"ai-introduction"`
						AstarSearch struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"astar-search"`
						AlphaBetaPruning struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"alpha-beta-pruning"`
						CombinatorialSearchTheory struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"combinatorial-search-theory"`
						RichmanGames struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"richman-games"`
						MachineLearning struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"machine-learning"`
						ImageAnalysis struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"image-analysis"`
						Nlp struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"nlp"`
						StatisticsFoundations struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"statistics-foundations"`
					} `json:"chapterDict"`
				} `json:"ai"`
				Cpp struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						CppIntroduction struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"cpp-introduction"`
						CppStrings struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"cpp-strings"`
						Classes struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"classes"`
						Stl struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"stl"`
						Inheritance struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"inheritance"`
						CppDebugging struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"cpp-debugging"`
						OtherConcepts struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"other-concepts"`
					} `json:"chapterDict"`
				} `json:"cpp"`
				Java struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						JavaIntroduction struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"java-introduction"`
						JavaStrings struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"java-strings"`
						Bignumber struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"bignumber"`
						JavaDataStructure struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"java-data-structure"`
						Oop struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"oop"`
						HandlingExceptions struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"handling-exceptions"`
						JavaAdvanced struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"java-advanced"`
					} `json:"chapterDict"`
				} `json:"java"`
				Python struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						PyIntroduction struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"py-introduction"`
						PyBasicDataTypes struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"py-basic-data-types"`
						PyStrings struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"py-strings"`
						PySets struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"py-sets"`
						PyMath struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"py-math"`
						PyItertools struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"py-itertools"`
						PyCollections struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"py-collections"`
						PyDateTime struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"py-date-time"`
						ErrorsExceptions struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"errors-exceptions"`
						PyClasses struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"py-classes"`
						PyBuiltIns struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"py-built-ins"`
						PyFunctionals struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"py-functionals"`
						PyRegex struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"py-regex"`
						XML struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"xml"`
						ClosuresAndDecorators struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"closures-and-decorators"`
						Numpy struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"numpy"`
						PyDebugging struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"py-debugging"`
					} `json:"chapterDict"`
				} `json:"python"`
				Ruby struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						RubyTutorials struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"ruby-tutorials"`
						ControlStructures struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"control-structures"`
						RubyArrays struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"ruby-arrays"`
						RubyEnumerables struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"ruby-enumerables"`
						RubyMethods struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"ruby-methods"`
						RubyStrings struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"ruby-strings"`
					} `json:"chapterDict"`
				} `json:"ruby"`
				Sql struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						Select struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"select"`
						AdvancedSelect struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"advanced-select"`
						Aggregation struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"aggregation"`
						Join struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"join"`
						AdvancedJoin struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"advanced-join"`
						AlternativeQueries struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"alternative-queries"`
					} `json:"chapterDict"`
				} `json:"sql"`
				Databases struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						RelationalAlgebra struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"relational-algebra"`
						Indexes struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"indexes"`
						Olap struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"olap"`
						SetAndAlgebra struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"set-and-algebra"`
						XpathQueries struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"xpath-queries"`
						DatabaseNormalization struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"database-normalization"`
					} `json:"chapterDict"`
				} `json:"databases"`
				DistributedSystems struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						DistributedMcq struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"distributed-mcq"`
						ClientServer struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"client-server"`
						MapreduceBasics struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"mapreduce-basics"`
					} `json:"chapterDict"`
				} `json:"distributed-systems"`
				Shell struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						Bash struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"bash"`
						Textpro struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"textpro"`
						ArraysInBash struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"arrays-in-bash"`
						GrepSedAwk struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"grep-sed-awk"`
					} `json:"chapterDict"`
				} `json:"shell"`
				Fp struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						Intro struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"intro"`
						FpRecursion struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"fp-recursion"`
						Ds struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"ds"`
						Dp struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"dp"`
						PersistentDs struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"persistent-ds"`
						Misc struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"misc"`
						Parsers struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"parsers"`
						Compilers struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"compilers"`
					} `json:"chapterDict"`
				} `json:"fp"`
				Regex struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						ReIntroduction struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"re-introduction"`
						ReCharacterClass struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"re-character-class"`
						ReRepetitions struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"re-repetitions"`
						GroupingAndCapturing struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"grouping-and-capturing"`
						Backreferences struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"backreferences"`
						Assertions struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"assertions"`
						ReApplications struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"re-applications"`
					} `json:"chapterDict"`
				} `json:"regex"`
				GeneralProgramming struct {
					ID          int           `json:"id"`
					Name        string        `json:"name"`
					Slug        string        `json:"slug"`
					Chapters    []interface{} `json:"chapters"`
					ChapterDict struct {
					} `json:"chapterDict"`
				} `json:"general-programming"`
				Security struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					Chapters []struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Hidden bool   `json:"hidden"`
					} `json:"chapters"`
					ChapterDict struct {
						Functions struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"functions"`
						Concepts struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"concepts"`
						Cryptography struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Slug   string `json:"slug"`
							Hidden bool   `json:"hidden"`
						} `json:"cryptography"`
					} `json:"chapterDict"`
				} `json:"security"`
			} `json:"dict"`
			DidInvalidate bool `json:"didInvalidate"`
		} `json:"domains"`
		Dashboard struct {
			DidInvalidate       bool          `json:"didInvalidate"`
			HasData             bool          `json:"hasData"`
			ShowAllTracks       bool          `json:"showAllTracks"`
			HasRecentTracks     bool          `json:"hasRecentTracks"`
			RecentTracks        []interface{} `json:"recentTracks"`
			TrackGroups         []interface{} `json:"trackGroups"`
			RecommendedTracks   []interface{} `json:"recommendedTracks"`
			RecommendedContests []interface{} `json:"recommendedContests"`
			MyTracks            []interface{} `json:"myTracks"`
		} `json:"dashboard"`
		Notifications struct {
			Summary struct {
			} `json:"summary"`
			List []interface{} `json:"list"`
		} `json:"notifications"`
		Messages struct {
			Threads []interface{} `json:"threads"`
		} `json:"messages"`
		Profile struct {
			ContestHistory struct {
				List  []interface{} `json:"list"`
				Total int           `json:"total"`
			} `json:"contestHistory"`
			Badges     []interface{} `json:"badges"`
			Additional struct {
			} `json:"additional"`
			BadgesVersion int `json:"badges_version"`
			Scores        struct {
			} `json:"scores"`
			Rating            []interface{} `json:"rating"`
			SubmissionHistory struct {
			} `json:"submissionHistory"`
			RecentChallenges struct {
				List     []interface{} `json:"list"`
				Cursor   interface{}   `json:"cursor"`
				LastPage bool          `json:"lastPage"`
			} `json:"recentChallenges"`
			RecentDiscussions []interface{} `json:"recentDiscussions"`
			ThreadID          interface{}   `json:"threadId"`
		} `json:"profile"`
		ViewProfiles struct {
		} `json:"viewProfiles"`
		Calendar struct {
			AllEvents struct {
				Events []interface{} `json:"events"`
			} `json:"allEvents"`
		} `json:"calendar"`
		Contests struct {
			Active struct {
				Total struct {
				} `json:"total"`
				Track struct {
				} `json:"track"`
				DidInvalidate bool `json:"didInvalidate"`
				PartialData   bool `json:"partialData"`
				LoadedFilter  struct {
				} `json:"loadedFilter"`
			} `json:"active"`
			AllContest struct {
				Contest struct {
					Master struct {
						ID                          int         `json:"id"`
						Name                        string      `json:"name"`
						Slug                        string      `json:"slug"`
						CreatedAt                   time.Time   `json:"created_at"`
						UpdatedAt                   time.Time   `json:"updated_at"`
						Starttime                   interface{} `json:"starttime"`
						Endtime                     interface{} `json:"endtime"`
						Timezone                    string      `json:"timezone"`
						Homepage                    string      `json:"homepage"`
						Tagline                     string      `json:"tagline"`
						Description                 string      `json:"description"`
						HomepageBackgroundColor     string      `json:"homepage_background_color"`
						Notification                interface{} `json:"notification"`
						TemplateID                  int         `json:"template_id"`
						ExposeStats                 interface{} `json:"expose_stats"`
						Public                      bool        `json:"public"`
						TeamEvent                   bool        `json:"team_event"`
						RatingCategory              interface{} `json:"rating_category"`
						IsRatingUpdated             bool        `json:"is_rating_updated"`
						LeaderboardBackend          interface{} `json:"leaderboard_backend"`
						LeaderboardFormat           string      `json:"leaderboard_format"`
						PrimaryTrackID              interface{} `json:"primary_track_id"`
						CollegePublic               bool        `json:"college_public"`
						Rated                       bool        `json:"rated"`
						IsMultiRound                bool        `json:"is_multi_round"`
						ParentContestID             interface{} `json:"parent_contest_id"`
						PrimaryTagID                interface{} `json:"primary_tag_id"`
						Started                     bool        `json:"started"`
						Ended                       bool        `json:"ended"`
						EpochEndtime                int         `json:"epoch_endtime"`
						EpochStarttime              int         `json:"epoch_starttime"`
						TimeLeft                    interface{} `json:"time_left"`
						HideDifficulty              interface{} `json:"hide_difficulty"`
						HasTracks                   bool        `json:"has_tracks"`
						Archived                    bool        `json:"archived"`
						LeaderboardType             string      `json:"leaderboard_type"`
						Kind                        string      `json:"kind"`
						LeaderboardFreezeTime       interface{} `json:"leaderboard_freeze_time"`
						ShowPenalty                 bool        `json:"show_penalty"`
						Track                       interface{} `json:"track"`
						HideNavigation              interface{} `json:"hide_navigation"`
						ContestBroadcast            interface{} `json:"contest_broadcast"`
						HideLeaderboard             interface{} `json:"hide_leaderboard"`
						HideSubmissions             interface{} `json:"hide_submissions"`
						LeaderboardOutOfSync        interface{} `json:"leaderboard_out_of_sync"`
						LeaderboardOutOfSyncMessage interface{} `json:"leaderboard_out_of_sync_message"`
						ChallengesCount             interface{} `json:"challenges_count"`
						ShowParticipantsInfo        interface{} `json:"show_participants_info"`
						CustomLeaderboardColumnName interface{} `json:"custom_leaderboard_column_name"`
						DisableForum                bool        `json:"disable_forum"`
						DisableFsi                  interface{} `json:"disable_fsi"`
						HasCodesprintRegPage        interface{} `json:"has_codesprint_reg_page"`
						Hidden                      interface{} `json:"hidden"`
						CommentLiveSync             interface{} `json:"comment_live_sync"`
						CompanyAssociatedContest    interface{} `json:"company_associated_contest"`
						LimitedParticipants         interface{} `json:"limited_participants"`
						LeaderboardBroadcastMessage interface{} `json:"leaderboard_broadcast_message"`
						QualificationRuleType       interface{} `json:"qualification_rule_type"`
						QualificationRuleValue      int         `json:"qualification_rule_value"`
						QualificationRuleMsg        interface{} `json:"qualification_rule_msg"`
						MigrationStatus             interface{} `json:"migration_status"`
						MigrationDisabled           interface{} `json:"migration_disabled"`
						TestersContest              interface{} `json:"testers_contest"`
						TimeLimitedContest          bool        `json:"time_limited_contest"`
						HackerTimelimit             interface{} `json:"hacker_timelimit"`
						SchoolLeaderboardEnabled    bool        `json:"school_leaderboard_enabled"`
						OrganizationType            interface{} `json:"organization_type"`
						OrganizationName            interface{} `json:"organization_name"`
						Categories                  []struct {
							ID       int    `json:"id"`
							Name     string `json:"name"`
							Slug     string `json:"slug"`
							Children []struct {
								ID                string      `json:"id"`
								Name              string      `json:"name"`
								Slug              string      `json:"slug"`
								Priority          int         `json:"priority"`
								OrderChallengeBy  int         `json:"order_challenge_by"`
								ChallengesPerPage interface{} `json:"challenges_per_page"`
							} `json:"children"`
						} `json:"categories"`
						PrimaryTags []struct {
							ID        int    `json:"id"`
							Slug      string `json:"slug"`
							Name      string `json:"name"`
							IsTrack   bool   `json:"is_track"`
							Weightage int    `json:"weightage"`
							Ancestry  string `json:"ancestry"`
							Priority  int    `json:"priority"`
						} `json:"primary_tags"`
						EffectiveTimeLeft     interface{} `json:"effective_time_left"`
						EffectiveEpochEndtime interface{} `json:"effective_epoch_endtime"`
						APITime               int64       `json:"apiTime"`
					} `json:"master"`
				} `json:"contest"`
			} `json:"allContest"`
			Archived struct {
				Total         int  `json:"total"`
				PartialData   bool `json:"partialData"`
				DidInvalidate bool `json:"didInvalidate"`
				Page          struct {
				} `json:"page"`
			} `json:"archived"`
			College struct {
				Total         int  `json:"total"`
				PartialData   bool `json:"partialData"`
				DidInvalidate bool `json:"didInvalidate"`
				Page          struct {
				} `json:"page"`
			} `json:"college"`
		} `json:"contests"`
		Leaderboard struct {
			List struct {
			} `json:"list"`
			Recommendations struct {
				DidInvalidate  bool `json:"didInvalidate"`
				Recommendation struct {
				} `json:"recommendation"`
			} `json:"recommendations"`
		} `json:"leaderboard"`
		Challenges struct {
			ChallengeList struct {
				ChallengePages struct {
				} `json:"challengePages"`
				CurrentTrack struct {
				} `json:"current_track"`
				AncestoryID   string `json:"ancestoryId"`
				DidInvalidate bool   `json:"didInvalidate"`
			} `json:"challengeList"`
			ChapterList struct {
				ChapterListObj struct {
				} `json:"chapterListObj"`
			} `json:"chapterList"`
			Challenge struct {
				Master2DArray struct {
					Detail struct {
						CanBeViewed             bool          `json:"can_be_viewed"`
						Dynamic                 bool          `json:"dynamic"`
						HasStarted              bool          `json:"has_started"`
						HasEnded                bool          `json:"has_ended"`
						CountdownTime           int           `json:"countdown_time"`
						RequirementsDescription interface{}   `json:"requirements_description"`
						MaxScore                int           `json:"max_score"`
						Active                  bool          `json:"active"`
						EpochStarttime          interface{}   `json:"epoch_starttime"`
						EpochEndtime            interface{}   `json:"epoch_endtime"`
						TimeLeft                interface{}   `json:"time_left"`
						Factor                  int           `json:"factor"`
						ExpertSolutionStatus    bool          `json:"expert_solution_status"`
						CustomTabs              interface{}   `json:"custom_tabs"`
						TotalCount              int           `json:"total_count"`
						SolvedCount             int           `json:"solved_count"`
						SuccessRatio            float64       `json:"success_ratio"`
						IsEditorialAvailable    bool          `json:"is_editorial_available"`
						IsSolutionUnlocked      bool          `json:"is_solution_unlocked"`
						ContestSlug             string        `json:"contest_slug"`
						Topics                  []interface{} `json:"topics"`
						UserScore               int           `json:"user_score"`
						ID                      int           `json:"id"`
						Slug                    string        `json:"slug"`
						Name                    string        `json:"name"`
						Status                  interface{}   `json:"status"`
						CreatedAt               time.Time     `json:"created_at"`
						UpdatedAt               time.Time     `json:"updated_at"`
						Kind                    string        `json:"kind"`
						Preview                 string        `json:"preview"`
						CodecheckerHandle       interface{}   `json:"codechecker_handle"`
						Category                string        `json:"category"`
						Deleted                 bool          `json:"deleted"`
						CompanyID               interface{}   `json:"company_id"`
						DifficultyScore         string        `json:"difficulty_score"`
						MigratedResource        string        `json:"migrated_resource"`
						Level                   int           `json:"level"`
						IsCustom                bool          `json:"is_custom"`
						PlayerCount             int           `json:"player_count"`
						CustomCheckerLanguage   string        `json:"custom_checker_language"`
						CheckerProgram          interface{}   `json:"checker_program"`
						JudgebotLanguage        interface{}   `json:"judgebot_language"`
						Judgebot                interface{}   `json:"judgebot"`
						Onboarding              interface{}   `json:"onboarding"`
						CompileAndTest          bool          `json:"compile_and_test"`
						Languages               []string      `json:"languages"`
						IsText                  bool          `json:"is_text"`
						Custom                  bool          `json:"custom"`
						Track                   struct {
							ID        int    `json:"id"`
							Name      string `json:"name"`
							Slug      string `json:"slug"`
							TrackID   int    `json:"track_id"`
							TrackName string `json:"track_name"`
							TrackSlug string `json:"track_slug"`
						} `json:"track"`
						CustomCase                  bool          `json:"custom_case"`
						SubmitDisabled              bool          `json:"submit_disabled"`
						PublicTestCases             bool          `json:"public_test_cases"`
						PublicSolutions             bool          `json:"public_solutions"`
						CanSolve                    bool          `json:"can_solve"`
						Company                     interface{}   `json:"company"`
						Difficulty                  float64       `json:"difficulty"`
						Color                       interface{}   `json:"color"`
						PrimaryContest              interface{}   `json:"primary_contest"`
						SolvedScore                 float64       `json:"solved_score"`
						AvailableTranslations       []interface{} `json:"available_translations"`
						LeaderboardOutOfSync        interface{}   `json:"leaderboard_out_of_sync"`
						LeaderboardOutOfSyncMessage interface{}   `json:"leaderboard_out_of_sync_message"`
						RequireUnlock               bool          `json:"require_unlock"`
						ShowSkeliton                interface{}   `json:"show_skeliton"`
						DifficultyName              string        `json:"difficulty_name"`
						DefaultLanguage             interface{}   `json:"default_language"`
						CTemplate                   string        `json:"c_template"`
						CTemplateHead               string        `json:"c_template_head"`
						CTemplateTail               string        `json:"c_template_tail"`
						CppTemplate                 string        `json:"cpp_template"`
						CppTemplateHead             string        `json:"cpp_template_head"`
						CppTemplateTail             string        `json:"cpp_template_tail"`
						JavaTemplate                string        `json:"java_template"`
						JavaTemplateHead            string        `json:"java_template_head"`
						JavaTemplateTail            string        `json:"java_template_tail"`
						CsharpTemplate              string        `json:"csharp_template"`
						CsharpTemplateHead          string        `json:"csharp_template_head"`
						CsharpTemplateTail          string        `json:"csharp_template_tail"`
						PhpTemplate                 string        `json:"php_template"`
						PhpTemplateHead             string        `json:"php_template_head"`
						PhpTemplateTail             string        `json:"php_template_tail"`
						RubyTemplate                string        `json:"ruby_template"`
						RubyTemplateHead            string        `json:"ruby_template_head"`
						RubyTemplateTail            string        `json:"ruby_template_tail"`
						PythonTemplate              string        `json:"python_template"`
						PythonTemplateHead          string        `json:"python_template_head"`
						PythonTemplateTail          string        `json:"python_template_tail"`
						PerlTemplate                string        `json:"perl_template"`
						PerlTemplateHead            string        `json:"perl_template_head"`
						PerlTemplateTail            string        `json:"perl_template_tail"`
						HaskellTemplate             string        `json:"haskell_template"`
						HaskellTemplateHead         string        `json:"haskell_template_head"`
						HaskellTemplateTail         string        `json:"haskell_template_tail"`
						ClojureTemplate             string        `json:"clojure_template"`
						ClojureTemplateHead         string        `json:"clojure_template_head"`
						ClojureTemplateTail         string        `json:"clojure_template_tail"`
						ScalaTemplate               string        `json:"scala_template"`
						ScalaTemplateHead           string        `json:"scala_template_head"`
						ScalaTemplateTail           string        `json:"scala_template_tail"`
						LuaTemplate                 string        `json:"lua_template"`
						LuaTemplateHead             string        `json:"lua_template_head"`
						LuaTemplateTail             string        `json:"lua_template_tail"`
						ErlangTemplate              string        `json:"erlang_template"`
						ErlangTemplateHead          string        `json:"erlang_template_head"`
						ErlangTemplateTail          string        `json:"erlang_template_tail"`
						JavascriptTemplate          string        `json:"javascript_template"`
						JavascriptTemplateHead      string        `json:"javascript_template_head"`
						JavascriptTemplateTail      string        `json:"javascript_template_tail"`
						GoTemplate                  string        `json:"go_template"`
						GoTemplateHead              string        `json:"go_template_head"`
						GoTemplateTail              string        `json:"go_template_tail"`
						Python3Template             string        `json:"python3_template"`
						Python3TemplateHead         string        `json:"python3_template_head"`
						Python3TemplateTail         string        `json:"python3_template_tail"`
						ObjectivecTemplate          string        `json:"objectivec_template"`
						ObjectivecTemplateHead      string        `json:"objectivec_template_head"`
						ObjectivecTemplateTail      string        `json:"objectivec_template_tail"`
						Java8Template               string        `json:"java8_template"`
						Java8TemplateHead           string        `json:"java8_template_head"`
						Java8TemplateTail           string        `json:"java8_template_tail"`
						SwiftTemplate               string        `json:"swift_template"`
						SwiftTemplateHead           string        `json:"swift_template_head"`
						SwiftTemplateTail           string        `json:"swift_template_tail"`
						Cpp14Template               string        `json:"cpp14_template"`
						Cpp14TemplateHead           string        `json:"cpp14_template_head"`
						Cpp14TemplateTail           string        `json:"cpp14_template_tail"`
						PypyTemplate                string        `json:"pypy_template"`
						PypyTemplateHead            string        `json:"pypy_template_head"`
						PypyTemplateTail            string        `json:"pypy_template_tail"`
						Pypy3Template               string        `json:"pypy3_template"`
						Pypy3TemplateHead           string        `json:"pypy3_template_head"`
						Pypy3TemplateTail           string        `json:"pypy3_template_tail"`
						KotlinTemplate              string        `json:"kotlin_template"`
						KotlinTemplateHead          string        `json:"kotlin_template_head"`
						KotlinTemplateTail          string        `json:"kotlin_template_tail"`
						Hacker                      struct {
							ID       int    `json:"id"`
							Username string `json:"username"`
							Avatar   string `json:"avatar"`
							IsAdmin  bool   `json:"is_admin"`
						} `json:"hacker"`
						BodyHTML              string `json:"body_html"`
						AuthorName            string `json:"author_name"`
						AuthorAvatar          string `json:"author_avatar"`
						IsPreviewContest      bool   `json:"is_preview_contest"`
						VisualOutputFlag      bool   `json:"visual_output_flag"`
						RealDynamic           bool   `json:"real_dynamic"`
						SubmittedHackersCount int    `json:"submitted_hackers_count"`
						HasAllTestCasesPublic bool   `json:"has_all_test_cases_public"`
						ModeratorOrSupport    bool   `json:"moderator_or_support"`
					} `json:"detail"`
					Rating struct {
					} `json:"rating"`
					Leaderboard struct {
					} `json:"leaderboard"`
					FileSubmissions []interface{} `json:"fileSubmissions"`
					Editorial       struct {
					} `json:"editorial"`
					BadgeTypes    []interface{} `json:"badgeTypes"`
					DidInvalidate bool          `json:"didInvalidate"`
				} `json:"master/2d-array"`
			} `json:"challenge"`
			Bookmarks struct {
				DidInvalidate bool `json:"didInvalidate"`
			} `json:"bookmarks"`
		} `json:"challenges"`
		Submissions struct {
			AllSubmissions struct {
			} `json:"allSubmissions"`
			ChallengeSubmissions struct {
			} `json:"challengeSubmissions"`
			GlobalStatus interface{} `json:"globalStatus"`
		} `json:"submissions"`
		Jobs struct {
			AllCompanies    []interface{} `json:"allCompanies"`
			ApplicationData struct {
			} `json:"applicationData"`
			CurrentCompany struct {
			} `json:"currentCompany"`
			CurrentJob struct {
			} `json:"currentJob"`
			AllConversations struct {
			} `json:"allConversations"`
			Search struct {
				Items   []interface{} `json:"items"`
				Filters struct {
				} `json:"filters"`
				Total        int  `json:"total"`
				Length       int  `json:"length"`
				CompanyLimit int  `json:"company_limit"`
				ByCompany    bool `json:"by_company"`
			} `json:"search"`
			States struct {
			} `json:"states"`
			Countries    []interface{} `json:"countries"`
			MessageCount struct {
				UnreadCount int `json:"unreadCount"`
			} `json:"messageCount"`
		} `json:"jobs"`
		CodeSnippets struct {
			CodeSnippet     interface{} `json:"codeSnippet"`
			CodeSnippetList struct {
				RevalidateData bool          `json:"_revalidate_data"`
				Total          interface{}   `json:"total"`
				Snippets       []interface{} `json:"snippets"`
			} `json:"codeSnippetList"`
			ForkedSnippets struct {
			} `json:"forkedSnippets"`
		} `json:"codeSnippets"`
		Feed struct {
			Feed       []interface{} `json:"feed"`
			Offset     int           `json:"offset"`
			KnownTotal interface{}   `json:"known_total"`
		} `json:"feed"`
		Badges struct {
			Badge struct {
			} `json:"badge"`
			TrackChapterBadgeMapping struct {
				Algorithms           []string `json:"algorithms"`
				DataStructures       []string `json:"data-structures"`
				Cpp                  []string `json:"cpp"`
				Java                 []string `json:"java"`
				Python               []string `json:"python"`
				Sql                  []string `json:"sql"`
				C                    []string `json:"c"`
				Ruby                 []string `json:"ruby"`
				Three0DaysOfCode     []string `json:"30-days-of-code"`
				One0DaysOfJavascript []string `json:"10-days-of-javascript"`
				One0DaysOfStatistics []string `json:"10-days-of-statistics"`
			} `json:"trackChapterBadgeMapping"`
			BadgeToTracksMap struct {
				ProblemSolving       []string `json:"problem-solving"`
				Cpp                  []string `json:"cpp"`
				Java                 []string `json:"java"`
				Python               []string `json:"python"`
				Sql                  []string `json:"sql"`
				C                    []string `json:"c"`
				Ruby                 []string `json:"ruby"`
				Three0DaysOfCode     []string `json:"30-days-of-code"`
				One0DaysOfJavascript []string `json:"10-days-of-javascript"`
				One0DaysOfStatistics []string `json:"10-days-of-statistics"`
			} `json:"badgeToTracksMap"`
		} `json:"badges"`
		Forum struct {
			Comments struct {
			} `json:"comments"`
			CommentsDict struct {
			} `json:"commentsDict"`
		} `json:"forum"`
		Playlist struct {
		} `json:"playlist"`
		NativeAds struct {
		} `json:"nativeAds"`
		UIState struct {
		} `json:"uiState"`
	} `json:"community"`
	Work struct {
	} `json:"work"`
	Metadata struct {
		Manifest struct {
			Badges10DaysOfJavascriptSvg                                                                           string `json:"badges/10-days-of-javascript.svg"`
			Badges10DaysOfStatisticsSvg                                                                           string `json:"badges/10-days-of-statistics.svg"`
			Badges30DaysOfCodeSvg                                                                                 string `json:"badges/30-days-of-code.svg"`
			BadgesAiSvg                                                                                           string `json:"badges/ai.svg"`
			BadgesAlgorithmsSvg                                                                                   string `json:"badges/algorithms.svg"`
			BadgesBadgeAlgoPng                                                                                    string `json:"badges/badge-algo.png"`
			BadgesBadgeFpPng                                                                                      string `json:"badges/badge-fp.png"`
			BadgesBadgeMlPng                                                                                      string `json:"badges/badge-ml.png"`
			BadgesBronzeSvg                                                                                       string `json:"badges/bronze.svg"`
			BadgesBronzeSmallSvg                                                                                  string `json:"badges/bronze_small.svg"`
			BadgesCSvg                                                                                            string `json:"badges/c.svg"`
			BadgesCppSvg                                                                                          string `json:"badges/cpp.svg"`
			BadgesCrackingTheCodingInterviewSvg                                                                   string `json:"badges/cracking-the-coding-interview.svg"`
			BadgesDataStructuresSvg                                                                               string `json:"badges/data-structures.svg"`
			BadgesDatabasesSvg                                                                                    string `json:"badges/databases.svg"`
			BadgesDistributedSystemsSvg                                                                           string `json:"badges/distributed-systems.svg"`
			BadgesFpSvg                                                                                           string `json:"badges/fp.svg"`
			BadgesGoldSvg                                                                                         string `json:"badges/gold.svg"`
			BadgesGoldSmallSvg                                                                                    string `json:"badges/gold_small.svg"`
			BadgesJavaSvg                                                                                         string `json:"badges/java.svg"`
			BadgesJavascriptSvg                                                                                   string `json:"badges/javascript.svg"`
			BadgesLinkedinPlacementsSvg                                                                           string `json:"badges/linkedin-placements.svg"`
			BadgesMathematicsAltSvg                                                                               string `json:"badges/mathematics-alt.svg"`
			BadgesMathematicsSvg                                                                                  string `json:"badges/mathematics.svg"`
			BadgesProblemSolvingSvg                                                                               string `json:"badges/problem-solving.svg"`
			BadgesPythonSvg                                                                                       string `json:"badges/python.svg"`
			BadgesRBadgeAlgoPng                                                                                   string `json:"badges/r-badge-algo.png"`
			BadgesRegexSvg                                                                                        string `json:"badges/regex.svg"`
			BadgesRubySvg                                                                                         string `json:"badges/ruby.svg"`
			BadgesRubyOldSvg                                                                                      string `json:"badges/ruby_old.svg"`
			BadgesSecuritySvg                                                                                     string `json:"badges/security.svg"`
			BadgesShellSvg                                                                                        string `json:"badges/shell.svg"`
			BadgesSilverSvg                                                                                       string `json:"badges/silver.svg"`
			BadgesSilverSmallSvg                                                                                  string `json:"badges/silver_small.svg"`
			BadgesSqlSvg                                                                                          string `json:"badges/sql.svg"`
			BadgesStarOutlinePng                                                                                  string `json:"badges/star-outline.png"`
			BadgesStarPng                                                                                         string `json:"badges/star.png"`
			BadgesStatisticsSvg                                                                                   string `json:"badges/statistics.svg"`
			BadgesTicTacToeOnePng                                                                                 string `json:"badges/tic-tac-toe-one.png"`
			BrandHMarkSmPng                                                                                       string `json:"brand/h_mark_sm.png"`
			BrandHMarkSmSvg                                                                                       string `json:"brand/h_mark_sm.svg"`
			BrandHackerRankWorkInverseSvg                                                                         string `json:"brand/hackerRank_work--inverse.svg"`
			BrandHrLogoNewBlackGreenSvg                                                                           string `json:"brand/hr-logo-new-black-green.svg"`
			BrandLogoNewWhiteGreenSvg                                                                             string `json:"brand/logo-new-white-green.svg"`
			BrandLogoWhiteGreenSvg                                                                                string `json:"brand/logo-white-green.svg"`
			BrandTypemark60X200Svg                                                                                string `json:"brand/typemark_60x200.svg"`
			ChallengesCompanyPagePng                                                                              string `json:"challenges/company_page.png"`
			ChallengesDbPng                                                                                       string `json:"challenges/db.png"`
			ChallengesDroidPng                                                                                    string `json:"challenges/droid.png"`
			ChallengesFilterIconSvg                                                                               string `json:"challenges/filter-icon.svg"`
			ChallengesGuyOnCompJpg                                                                                string `json:"challenges/guy-on-comp.jpg"`
			ChallengesInsightsPng                                                                                 string `json:"challenges/insights.png"`
			ChallengesLibraryPng                                                                                  string `json:"challenges/library.png"`
			ChallengesLockedContentSvg                                                                            string `json:"challenges/locked-content.svg"`
			ChallengesSidebarDiscussPng                                                                           string `json:"challenges/sidebar_discuss.png"`
			ChallengesSidebarTopscorersPng                                                                        string `json:"challenges/sidebar_topscorers.png"`
			ChallengesSidebarTutorialPng                                                                          string `json:"challenges/sidebar_tutorial.png"`
			ChallengesSudoPng                                                                                     string `json:"challenges/sudo.png"`
			ChallengesVmwareGuyPng                                                                                string `json:"challenges/vmware_guy.png"`
			Dashboard10DaysOfJavascriptSvg                                                                        string `json:"dashboard/10-days-of-javascript.svg"`
			Dashboard10DaysOfStatisticsSvg                                                                        string `json:"dashboard/10-days-of-statistics.svg"`
			Dashboard101HackPng                                                                                   string `json:"dashboard/101hack.png"`
			Dashboard30DaysOfCodeSvg                                                                              string `json:"dashboard/30-days-of-code.svg"`
			Dashboard30DaysofCode1Svg                                                                             string `json:"dashboard/30daysofCode1.svg"`
			Dashboard30DaysofCode2Svg                                                                             string `json:"dashboard/30daysofCode2.svg"`
			Dashboard30DaysofCode3Svg                                                                             string `json:"dashboard/30daysofCode3.svg"`
			DashboardAiSvg                                                                                        string `json:"dashboard/ai.svg"`
			DashboardAlgorithmsSvg                                                                                string `json:"dashboard/algorithms.svg"`
			DashboardCSvg                                                                                         string `json:"dashboard/c.svg"`
			DashboardCppSvg                                                                                       string `json:"dashboard/cpp.svg"`
			DashboardCrackingTheCodingInterviewSvg                                                                string `json:"dashboard/cracking-the-coding-interview.svg"`
			DashboardCtci1Svg                                                                                     string `json:"dashboard/ctci1.svg"`
			DashboardCtci2Svg                                                                                     string `json:"dashboard/ctci2.svg"`
			DashboardCtci3Svg                                                                                     string `json:"dashboard/ctci3.svg"`
			DashboardDataStructuresSvg                                                                            string `json:"dashboard/data-structures.svg"`
			DashboardDatabasesSvg                                                                                 string `json:"dashboard/databases.svg"`
			DashboardDefaultPng                                                                                   string `json:"dashboard/default.png"`
			DashboardDistributedSystemsSvg                                                                        string `json:"dashboard/distributed-systems.svg"`
			DashboardFpSvg                                                                                        string `json:"dashboard/fp.svg"`
			DashboardGeneralProgrammingPng                                                                        string `json:"dashboard/general-programming.png"`
			DashboardGeneralProgrammingSvg                                                                        string `json:"dashboard/general-programming.svg"`
			DashboardHourrankPng                                                                                  string `json:"dashboard/hourrank.png"`
			DashboardInfinitumPng                                                                                 string `json:"dashboard/infinitum.png"`
			DashboardJavaSvg                                                                                      string `json:"dashboard/java.svg"`
			DashboardLinkedinPlacementsSvg                                                                        string `json:"dashboard/linkedin-placements.svg"`
			DashboardMathSvg                                                                                      string `json:"dashboard/math.svg"`
			DashboardMathematicsSvg                                                                               string `json:"dashboard/mathematics.svg"`
			DashboardOnboarding30DaysofCode1Svg                                                                   string `json:"dashboard/onboarding/30daysofCode1.svg"`
			DashboardOnboarding30DaysofCode2Svg                                                                   string `json:"dashboard/onboarding/30daysofCode2.svg"`
			DashboardOnboarding30DaysofCode3Svg                                                                   string `json:"dashboard/onboarding/30daysofCode3.svg"`
			DashboardOnboardingCtci1Svg                                                                           string `json:"dashboard/onboarding/ctci1.svg"`
			DashboardOnboardingCtci2Svg                                                                           string `json:"dashboard/onboarding/ctci2.svg"`
			DashboardOnboardingCtci3Svg                                                                           string `json:"dashboard/onboarding/ctci3.svg"`
			DashboardPythonSvg                                                                                    string `json:"dashboard/python.svg"`
			DashboardRegexSvg                                                                                     string `json:"dashboard/regex.svg"`
			DashboardRookiePng                                                                                    string `json:"dashboard/rookie.png"`
			DashboardRubySvg                                                                                      string `json:"dashboard/ruby.svg"`
			DashboardSecuritySvg                                                                                  string `json:"dashboard/security.svg"`
			DashboardShellSvg                                                                                     string `json:"dashboard/shell.svg"`
			DashboardSqlSvg                                                                                       string `json:"dashboard/sql.svg"`
			DashboardSurveyFocusJobSvg                                                                            string `json:"dashboard/survey/focus-job.svg"`
			DashboardSurveyFocusLearnSvg                                                                          string `json:"dashboard/survey/focus-learn.svg"`
			DashboardSurveyFocusParticipateSvg                                                                    string `json:"dashboard/survey/focus-participate.svg"`
			DashboardSurveyFocusPracticeSvg                                                                       string `json:"dashboard/survey/focus-practice.svg"`
			DashboardSurveyFocusPrepSvg                                                                           string `json:"dashboard/survey/focus-prep.svg"`
			DashboardSurveyProSelectionSvg                                                                        string `json:"dashboard/survey/pro-selection.svg"`
			DashboardSurveyStudentSelectionSvg                                                                    string `json:"dashboard/survey/student-selection.svg"`
			DashboardUniversityPng                                                                                string `json:"dashboard/university.png"`
			DashboardWeeklyPng                                                                                    string `json:"dashboard/weekly.png"`
			DashboardWomensCodesprintPng                                                                          string `json:"dashboard/womens-codesprint.png"`
			DashboardWorldCodesprintPng                                                                           string `json:"dashboard/world-codesprint.png"`
			FacebookColoredSvg                                                                                    string `json:"facebook-colored.svg"`
			FaviconIco                                                                                            string `json:"favicon.ico"`
			FaviconPng                                                                                            string `json:"favicon.png"`
			FbSharePng                                                                                            string `json:"fb-share.png"`
			FeedbackFeedbackBannerDefaultSvg                                                                      string `json:"feedback/feedback-banner-default.svg"`
			FeedbackFeedbackBannerLibrarySvg                                                                      string `json:"feedback/feedback-banner-library.svg"`
			FeedbackFeedbackBannerTestsSvg                                                                        string `json:"feedback/feedback-banner-tests.svg"`
			FeedbackFeedbackCantsaySvg                                                                            string `json:"feedback/feedback-cantsay.svg"`
			FeedbackFeedbackLikeSvg                                                                               string `json:"feedback/feedback-like.svg"`
			FeedbackFeedbackNoSvg                                                                                 string `json:"feedback/feedback-no.svg"`
			ForumMailerFollowIconSmallPng                                                                         string `json:"forum_mailer/follow-icon-small.png"`
			ForumMailerNotifymarkerPng                                                                            string `json:"forum_mailer/notifymarker.png"`
			FourohfourPng                                                                                         string `json:"fourohfour.png"`
			GithubColoredSvg                                                                                      string `json:"github-colored.svg"`
			GoogleColoredSvg                                                                                      string `json:"google-colored.svg"`
			IconsPartyPng                                                                                         string `json:"icons/party.png"`
			JobsCustomizedmatchingPng                                                                             string `json:"jobs/customizedmatching.png"`
			JobsDefaultlogoPng                                                                                    string `json:"jobs/defaultlogo.png"`
			JobsLappyPng                                                                                          string `json:"jobs/lappy.png"`
			LinkedinColoredSvg                                                                                    string `json:"linkedin-colored.svg"`
			NavDotsPng                                                                                            string `json:"nav_dots.png"`
			OnboardingFauxLeaderboardJpg                                                                          string `json:"onboarding/faux_leaderboard.jpg"`
			OnboardingHrMascotErrorPng                                                                            string `json:"onboarding/hr-mascot-error.png"`
			OnboardingHrMascotPng                                                                                 string `json:"onboarding/hr-mascot.png"`
			ProfileDefaultAvatarBgPng                                                                             string `json:"profile/default_avatar_bg.png"`
			SvgsHexHollowSvg                                                                                      string `json:"svgs/hex_hollow.svg"`
			SvgsMegaPhoneSvg                                                                                      string `json:"svgs/mega-phone.svg"`
			SvgsNewBannerSvg                                                                                      string `json:"svgs/new-banner.svg"`
			SvgsNoBadgesSvg                                                                                       string `json:"svgs/no-badges.svg"`
			SvgsNoJobsIconSvg                                                                                     string `json:"svgs/no-jobs-icon.svg"`
			TutorialsBgPng                                                                                        string `json:"tutorials/bg.png"`
			TutorialsCalendarPng                                                                                  string `json:"tutorials/calendar.png"`
			TutorialsDataPng                                                                                      string `json:"tutorials/data.png"`
			TutorialsDicePng                                                                                      string `json:"tutorials/dice.png"`
			TutorialsGaylePhotoPng                                                                                string `json:"tutorials/gayle-photo.png"`
			TutorialsGraphPng                                                                                     string `json:"tutorials/graph.png"`
			TutorialsInterviewIconPng                                                                             string `json:"tutorials/interview-icon.png"`
			TutorialsKathrynPicPng                                                                                string `json:"tutorials/kathryn_pic.png"`
			TutorialsLiPracticePhotoPng                                                                           string `json:"tutorials/li-practice-photo.png"`
			TutorialsPatternPng                                                                                   string `json:"tutorials/pattern.png"`
			TutorialsPolishIconSvg                                                                                string `json:"tutorials/polish-icon.svg"`
			TutorialsReadIconSvg                                                                                  string `json:"tutorials/read-icon.svg"`
			TutorialsSolveIconSvg                                                                                 string `json:"tutorials/solve-icon.svg"`
			TutorialsStatsPicPng                                                                                  string `json:"tutorials/stats_pic.png"`
			TutorialsTagPng                                                                                       string `json:"tutorials/tag.png"`
			TutorialsUnlockPng                                                                                    string `json:"tutorials/unlock.png"`
			TutorialsVerifyIconPng                                                                                string `json:"tutorials/verify-icon.png"`
			TutorialsVideoIconPng                                                                                 string `json:"tutorials/video-icon.png"`
			TweetFillerPng                                                                                        string `json:"tweet-filler.png"`
			WorkInviteEmptyTemplatesSvg                                                                           string `json:"work/invite/empty-templates.svg"`
			WorkLibraryEmptyLibrarySvg                                                                            string `json:"work/library/empty-library.svg"`
			WorkQuestionsRocketShipHeroSvg                                                                        string `json:"work/questions/rocket-ship_hero.svg"`
			WorkRbaCandidateEvaluationSvg                                                                         string `json:"work/rba/candidate_evaluation.svg"`
			WorkRbaOnlineIDESvg                                                                                   string `json:"work/rba/online_IDE.svg"`
			WorkRbaRealWorldAssesmentsSvg                                                                         string `json:"work/rba/real_world_assesments.svg"`
			WorkRbaRoleSpecificSvg                                                                                string `json:"work/rba/role_specific.svg"`
			WorkTeamsCompanyAdminSvg                                                                              string `json:"work/teams/company_admin.svg"`
			WorkTeamsDeveloperSvg                                                                                 string `json:"work/teams/developer.svg"`
			WorkTeamsNoTeamSvg                                                                                    string `json:"work/teams/no-team.svg"`
			WorkTeamsNoUserSvg                                                                                    string `json:"work/teams/no-user.svg"`
			WorkTeamsRecruiterSvg                                                                                 string `json:"work/teams/recruiter.svg"`
			WorkTeamsTeamAdminSvg                                                                                 string `json:"work/teams/team_admin.svg"`
			WorkTestsAddQuestionFromLibrarySvg                                                                    string `json:"work/tests/add-question-from-library.svg"`
			WorkTestsAlertSmallSvg                                                                                string `json:"work/tests/alert-small.svg"`
			WorkTestsCreateOwnQuestionSvg                                                                         string `json:"work/tests/create-own-question.svg"`
			WorkTestsEmptyStateSvg                                                                                string `json:"work/tests/empty-state.svg"`
			WorkTestsEmptyTestsSvg                                                                                string `json:"work/tests/empty-tests.svg"`
			WorkTestsIcnBlankSvg                                                                                  string `json:"work/tests/icn-blank.svg"`
			WorkTestsIcnTemplateSvg                                                                               string `json:"work/tests/icn-template.svg"`
			WorkTestsInviteIconSvg                                                                                string `json:"work/tests/invite-icon.svg"`
			WorkTestsKudosIconSvg                                                                                 string `json:"work/tests/kudos-icon.svg"`
			WorkTestsReportsEmptyStateSvg                                                                         string `json:"work/tests/reports/empty-state.svg"`
			WorkTestsReportsInviteIconSvg                                                                         string `json:"work/tests/reports/invite-icon.svg"`
			WorkTestsReportsKudosIconSvg                                                                          string `json:"work/tests/reports/kudos-icon.svg"`
			WorkTestsSurveySvg                                                                                    string `json:"work/tests/survey.svg"`
			WorkTestsWelcomeScreenSvg                                                                             string `json:"work/tests/welcome_screen.svg"`
			UIKitMinCSS                                                                                           string `json:"ui-kit.min.css"`
			UIKitMinJs                                                                                            string `json:"ui-kit.min.js"`
			BackboneStylesDashboardCSS                                                                            string `json:"backbone_styles/dashboard.css"`
			BackboneStylesHackerrankCoreCSS                                                                       string `json:"backbone_styles/hackerrank-core.css"`
			BackboneStylesHackerrankLibrariesCSS                                                                  string `json:"backbone_styles/hackerrank_libraries.css"`
			CkeditorCustomCSS                                                                                     string `json:"ckeditor-custom.css"`
			CoreHelperLibScss                                                                                     string `json:"core/helper-lib.scss"`
			CoreHelperScss                                                                                        string `json:"core/helper.scss"`
			CoreMixinsScss                                                                                        string `json:"core/mixins.scss"`
			CoreResetScss                                                                                         string `json:"core/reset.scss"`
			CoreTypographyLibScss                                                                                 string `json:"core/typography-lib.scss"`
			CoreVariablesScss                                                                                     string `json:"core/variables.scss"`
			FontIconsScss                                                                                         string `json:"font_icons.scss"`
			HackdownScss                                                                                          string `json:"hackdown.scss"`
			ModulesCodeeditorScss                                                                                 string `json:"modules/codeeditor.scss"`
			OverrideJobsMessageCkeditorCSS                                                                        string `json:"override_jobs_message_ckeditor.css"`
			ReactCkeditorCSS                                                                                      string `json:"react-ckeditor.css"`
			ThemeMCkeditorCSS                                                                                     string `json:"theme_m_ckeditor.css"`
			OpenSansBoldWoff                                                                                      string `json:"open-sans-bold.woff"`
			OpenSansBoldWoff2                                                                                     string `json:"open-sans-bold.woff2"`
			OpenSansWoff                                                                                          string `json:"open-sans.woff"`
			OpenSansWoff2                                                                                         string `json:"open-sans.woff2"`
			SourceCodeProBoldWoff                                                                                 string `json:"source-code-pro-bold.woff"`
			SourceCodeProBoldWoff2                                                                                string `json:"source-code-pro-bold.woff2"`
			SourceCodeProWoff                                                                                     string `json:"source-code-pro.woff"`
			SourceCodeProWoff2                                                                                    string `json:"source-code-pro.woff2"`
			UIKitIconEot                                                                                          string `json:"UIKitIcon.eot"`
			UIKitIconSvg                                                                                          string `json:"UIKitIcon.svg"`
			UIKitIconTtf                                                                                          string `json:"UIKitIcon.ttf"`
			UIKitIconWoff                                                                                         string `json:"UIKitIcon.woff"`
			UIKitIconWoff2                                                                                        string `json:"UIKitIcon.woff2"`
			UIIconScss                                                                                            string `json:"ui_icon.scss"`
			UIIconsSvgsActivitySvg                                                                                string `json:"ui_icons_svgs/activity.svg"`
			UIIconsSvgsAddUserSvg                                                                                 string `json:"ui_icons_svgs/add-user.svg"`
			UIIconsSvgsAlertTriangleSvg                                                                           string `json:"ui_icons_svgs/alert-triangle.svg"`
			UIIconsSvgsArrowLeftSvg                                                                               string `json:"ui_icons_svgs/arrow-left.svg"`
			UIIconsSvgsArrowRightSvg                                                                              string `json:"ui_icons_svgs/arrow-right.svg"`
			UIIconsSvgsBulbSvg                                                                                    string `json:"ui_icons_svgs/bulb.svg"`
			UIIconsSvgsCashSvg                                                                                    string `json:"ui_icons_svgs/cash.svg"`
			UIIconsSvgsCheckCircleSvg                                                                             string `json:"ui_icons_svgs/check-circle.svg"`
			UIIconsSvgsCheckSquareSvg                                                                             string `json:"ui_icons_svgs/check-square.svg"`
			UIIconsSvgsChevronDownSvg                                                                             string `json:"ui_icons_svgs/chevron-down.svg"`
			UIIconsSvgsChevronRightSvg                                                                            string `json:"ui_icons_svgs/chevron-right.svg"`
			UIIconsSvgsCircleFilledSvg                                                                            string `json:"ui_icons_svgs/circle-filled.svg"`
			UIIconsSvgsCircleSvg                                                                                  string `json:"ui_icons_svgs/circle.svg"`
			UIIconsSvgsClipboardSvg                                                                               string `json:"ui_icons_svgs/clipboard.svg"`
			UIIconsSvgsCodeTrySvg                                                                                 string `json:"ui_icons_svgs/code-try.svg"`
			UIIconsSvgsCodeSvg                                                                                    string `json:"ui_icons_svgs/code.svg"`
			UIIconsSvgsCodepenSvg                                                                                 string `json:"ui_icons_svgs/codepen.svg"`
			UIIconsSvgsCompanyAdminSvg                                                                            string `json:"ui_icons_svgs/company-admin.svg"`
			UIIconsSvgsCrossCircleSvg                                                                             string `json:"ui_icons_svgs/cross-circle.svg"`
			UIIconsSvgsCrossSvg                                                                                   string `json:"ui_icons_svgs/cross.svg"`
			UIIconsSvgsDeleteSvg                                                                                  string `json:"ui_icons_svgs/delete.svg"`
			UIIconsSvgsDiscussionSvg                                                                              string `json:"ui_icons_svgs/discussion.svg"`
			UIIconsSvgsDollarSvg                                                                                  string `json:"ui_icons_svgs/dollar.svg"`
			UIIconsSvgsDownloadSvg                                                                                string `json:"ui_icons_svgs/download.svg"`
			UIIconsSvgsDragHandleSvg                                                                              string `json:"ui_icons_svgs/drag-handle.svg"`
			UIIconsSvgsEditPlainSvg                                                                               string `json:"ui_icons_svgs/edit-plain.svg"`
			UIIconsSvgsEditSvg                                                                                    string `json:"ui_icons_svgs/edit.svg"`
			UIIconsSvgsEditorialSvg                                                                               string `json:"ui_icons_svgs/editorial.svg"`
			UIIconsSvgsEmailSvg                                                                                   string `json:"ui_icons_svgs/email.svg"`
			UIIconsSvgsEmploymentSvg                                                                              string `json:"ui_icons_svgs/employment.svg"`
			UIIconsSvgsErrorSvg                                                                                   string `json:"ui_icons_svgs/error.svg"`
			UIIconsSvgsFacebookFilledSvg                                                                          string `json:"ui_icons_svgs/facebook-filled.svg"`
			UIIconsSvgsFacebookLineSvg                                                                            string `json:"ui_icons_svgs/facebook-line.svg"`
			UIIconsSvgsFacebookSelectedSvg                                                                        string `json:"ui_icons_svgs/facebook-selected.svg"`
			UIIconsSvgsFacebookSvg                                                                                string `json:"ui_icons_svgs/facebook.svg"`
			UIIconsSvgsFeatureFeedbackSvg                                                                         string `json:"ui_icons_svgs/feature-feedback.svg"`
			UIIconsSvgsFilterSvg                                                                                  string `json:"ui_icons_svgs/filter.svg"`
			UIIconsSvgsForkSvg                                                                                    string `json:"ui_icons_svgs/fork.svg"`
			UIIconsSvgsGithubSelectedSvg                                                                          string `json:"ui_icons_svgs/github-selected.svg"`
			UIIconsSvgsGithubSvg                                                                                  string `json:"ui_icons_svgs/github.svg"`
			UIIconsSvgsGlobeSvg                                                                                   string `json:"ui_icons_svgs/globe.svg"`
			UIIconsSvgsGraduationSvg                                                                              string `json:"ui_icons_svgs/graduation.svg"`
			UIIconsSvgsHappyFaceSvg                                                                               string `json:"ui_icons_svgs/happy-face.svg"`
			UIIconsSvgsHelpSvg                                                                                    string `json:"ui_icons_svgs/help.svg"`
			UIIconsSvgsHexagonFilledSvg                                                                           string `json:"ui_icons_svgs/hexagon-filled.svg"`
			UIIconsSvgsHexagonLineSvg                                                                             string `json:"ui_icons_svgs/hexagon-line.svg"`
			UIIconsSvgsHideSvg                                                                                    string `json:"ui_icons_svgs/hide.svg"`
			UIIconsSvgsHintSvg                                                                                    string `json:"ui_icons_svgs/hint.svg"`
			UIIconsSvgsHistorySvg                                                                                 string `json:"ui_icons_svgs/history.svg"`
			UIIconsSvgsHomeSvg                                                                                    string `json:"ui_icons_svgs/home.svg"`
			UIIconsSvgsInterviewSvg                                                                               string `json:"ui_icons_svgs/interview.svg"`
			UIIconsSvgsLinkedinFiledSvg                                                                           string `json:"ui_icons_svgs/linkedin-filed.svg"`
			UIIconsSvgsLinkedinFilledSvg                                                                          string `json:"ui_icons_svgs/linkedin-filled.svg"`
			UIIconsSvgsLinkedinLineSvg                                                                            string `json:"ui_icons_svgs/linkedin-line.svg"`
			UIIconsSvgsLinkedinSelectedSvg                                                                        string `json:"ui_icons_svgs/linkedin-selected.svg"`
			UIIconsSvgsLinkedinSvg                                                                                string `json:"ui_icons_svgs/linkedin.svg"`
			UIIconsSvgsLoadingSvg                                                                                 string `json:"ui_icons_svgs/loading.svg"`
			UIIconsSvgsLocationSvg                                                                                string `json:"ui_icons_svgs/location.svg"`
			UIIconsSvgsLockSvg                                                                                    string `json:"ui_icons_svgs/lock.svg"`
			UIIconsSvgsMailAlertSvg                                                                               string `json:"ui_icons_svgs/mail-alert.svg"`
			UIIconsSvgsMailSvg                                                                                    string `json:"ui_icons_svgs/mail.svg"`
			UIIconsSvgsMaximizeSvg                                                                                string `json:"ui_icons_svgs/maximize.svg"`
			UIIconsSvgsMenuMoreSvg                                                                                string `json:"ui_icons_svgs/menu-more.svg"`
			UIIconsSvgsMessageSvg                                                                                 string `json:"ui_icons_svgs/message.svg"`
			UIIconsSvgsMinimizeSvg                                                                                string `json:"ui_icons_svgs/minimize.svg"`
			UIIconsSvgsMinusCircleSvg                                                                             string `json:"ui_icons_svgs/minus-circle.svg"`
			UIIconsSvgsMonitorSvg                                                                                 string `json:"ui_icons_svgs/monitor.svg"`
			UIIconsSvgsNotificationSvg                                                                            string `json:"ui_icons_svgs/notification.svg"`
			UIIconsSvgsOpenNewWindowSvg                                                                           string `json:"ui_icons_svgs/open-new-window.svg"`
			UIIconsSvgsPlagiarismSvg                                                                              string `json:"ui_icons_svgs/plagiarism.svg"`
			UIIconsSvgsPlaySvg                                                                                    string `json:"ui_icons_svgs/play.svg"`
			UIIconsSvgsRatingsSvg                                                                                 string `json:"ui_icons_svgs/ratings.svg"`
			UIIconsSvgsResendSvg                                                                                  string `json:"ui_icons_svgs/resend.svg"`
			UIIconsSvgsResetCodeSvg                                                                               string `json:"ui_icons_svgs/reset-code.svg"`
			UIIconsSvgsRestoreSizeSvg                                                                             string `json:"ui_icons_svgs/restore-size.svg"`
			UIIconsSvgsSadFaceSvg                                                                                 string `json:"ui_icons_svgs/sad-face.svg"`
			UIIconsSvgsScoreSvg                                                                                   string `json:"ui_icons_svgs/score.svg"`
			UIIconsSvgsSearchSvg                                                                                  string `json:"ui_icons_svgs/search.svg"`
			UIIconsSvgsSettingsSvg                                                                                string `json:"ui_icons_svgs/settings.svg"`
			UIIconsSvgsShareSvg                                                                                   string `json:"ui_icons_svgs/share.svg"`
			UIIconsSvgsShowSvg                                                                                    string `json:"ui_icons_svgs/show.svg"`
			UIIconsSvgsSmileySvg                                                                                  string `json:"ui_icons_svgs/smiley.svg"`
			UIIconsSvgsStarFilledSvg                                                                              string `json:"ui_icons_svgs/star-filled.svg"`
			UIIconsSvgsStarSvg                                                                                    string `json:"ui_icons_svgs/star.svg"`
			UIIconsSvgsSubmissionsSvg                                                                             string `json:"ui_icons_svgs/submissions.svg"`
			UIIconsSvgsSuccessSvg                                                                                 string `json:"ui_icons_svgs/success.svg"`
			UIIconsSvgsTeamAdminSvg                                                                               string `json:"ui_icons_svgs/team-admin.svg"`
			UIIconsSvgsThumbsDown2Svg                                                                             string `json:"ui_icons_svgs/thumbs-down-2.svg"`
			UIIconsSvgsThumbsDownFilledSvg                                                                        string `json:"ui_icons_svgs/thumbs-down-filled.svg"`
			UIIconsSvgsThumbsDownSvg                                                                              string `json:"ui_icons_svgs/thumbs-down.svg"`
			UIIconsSvgsThumbsUp2Svg                                                                               string `json:"ui_icons_svgs/thumbs-up-2.svg"`
			UIIconsSvgsThumbsUpFilledSvg                                                                          string `json:"ui_icons_svgs/thumbs-up-filled.svg"`
			UIIconsSvgsThumbsUpLightSvg                                                                           string `json:"ui_icons_svgs/thumbs-up-light.svg"`
			UIIconsSvgsThumbsUpSvg                                                                                string `json:"ui_icons_svgs/thumbs-up.svg"`
			UIIconsSvgsTimeSvg                                                                                    string `json:"ui_icons_svgs/time.svg"`
			UIIconsSvgsTrashSvg                                                                                   string `json:"ui_icons_svgs/trash.svg"`
			UIIconsSvgsTrophySvg                                                                                  string `json:"ui_icons_svgs/trophy.svg"`
			UIIconsSvgsTutorialSvg                                                                                string `json:"ui_icons_svgs/tutorial.svg"`
			UIIconsSvgsTwitterSvg                                                                                 string `json:"ui_icons_svgs/twitter.svg"`
			UIIconsSvgsUploadSvg                                                                                  string `json:"ui_icons_svgs/upload.svg"`
			UIIconsSvgsUserSvg                                                                                    string `json:"ui_icons_svgs/user.svg"`
			UIIconsSvgsWarningHexagonSvg                                                                          string `json:"ui_icons_svgs/warning-hexagon.svg"`
			CodemirrorModeAplAplJs                                                                                string `json:"codemirror/mode/apl/apl.js"`
			CodemirrorModeAsciiarmorAsciiarmorJs                                                                  string `json:"codemirror/mode/asciiarmor/asciiarmor.js"`
			CodemirrorModeAsn1Asn1Js                                                                              string `json:"codemirror/mode/asn.1/asn.1.js"`
			CodemirrorModeAsteriskAsteriskJs                                                                      string `json:"codemirror/mode/asterisk/asterisk.js"`
			CodemirrorModeBrainfuckBrainfuckJs                                                                    string `json:"codemirror/mode/brainfuck/brainfuck.js"`
			CodemirrorModeClikeClikeJs                                                                            string `json:"codemirror/mode/clike/clike.js"`
			CodemirrorModeClojureClojureJs                                                                        string `json:"codemirror/mode/clojure/clojure.js"`
			CodemirrorModeCmakeCmakeJs                                                                            string `json:"codemirror/mode/cmake/cmake.js"`
			CodemirrorModeCobolCobolJs                                                                            string `json:"codemirror/mode/cobol/cobol.js"`
			CodemirrorModeCoffeescriptCoffeescriptJs                                                              string `json:"codemirror/mode/coffeescript/coffeescript.js"`
			CodemirrorModeCommonlispCommonlispJs                                                                  string `json:"codemirror/mode/commonlisp/commonlisp.js"`
			CodemirrorModeCrystalCrystalJs                                                                        string `json:"codemirror/mode/crystal/crystal.js"`
			CodemirrorModeCSSCSSJs                                                                                string `json:"codemirror/mode/css/css.js"`
			CodemirrorModeCypherCypherJs                                                                          string `json:"codemirror/mode/cypher/cypher.js"`
			CodemirrorModeDDJs                                                                                    string `json:"codemirror/mode/d/d.js"`
			CodemirrorModeDartDartJs                                                                              string `json:"codemirror/mode/dart/dart.js"`
			CodemirrorModeDiffDiffJs                                                                              string `json:"codemirror/mode/diff/diff.js"`
			CodemirrorModeDjangoDjangoJs                                                                          string `json:"codemirror/mode/django/django.js"`
			CodemirrorModeDockerfileDockerfileJs                                                                  string `json:"codemirror/mode/dockerfile/dockerfile.js"`
			CodemirrorModeDtdDtdJs                                                                                string `json:"codemirror/mode/dtd/dtd.js"`
			CodemirrorModeDylanDylanJs                                                                            string `json:"codemirror/mode/dylan/dylan.js"`
			CodemirrorModeEbnfEbnfJs                                                                              string `json:"codemirror/mode/ebnf/ebnf.js"`
			CodemirrorModeEclEclJs                                                                                string `json:"codemirror/mode/ecl/ecl.js"`
			CodemirrorModeEiffelEiffelJs                                                                          string `json:"codemirror/mode/eiffel/eiffel.js"`
			CodemirrorModeElmElmJs                                                                                string `json:"codemirror/mode/elm/elm.js"`
			CodemirrorModeErlangErlangJs                                                                          string `json:"codemirror/mode/erlang/erlang.js"`
			CodemirrorModeFactorFactorJs                                                                          string `json:"codemirror/mode/factor/factor.js"`
			CodemirrorModeFclFclJs                                                                                string `json:"codemirror/mode/fcl/fcl.js"`
			CodemirrorModeForthForthJs                                                                            string `json:"codemirror/mode/forth/forth.js"`
			CodemirrorModeFortranFortranJs                                                                        string `json:"codemirror/mode/fortran/fortran.js"`
			CodemirrorModeGasGasJs                                                                                string `json:"codemirror/mode/gas/gas.js"`
			CodemirrorModeGfmGfmJs                                                                                string `json:"codemirror/mode/gfm/gfm.js"`
			CodemirrorModeGherkinGherkinJs                                                                        string `json:"codemirror/mode/gherkin/gherkin.js"`
			CodemirrorModeGoGoJs                                                                                  string `json:"codemirror/mode/go/go.js"`
			CodemirrorModeGroovyGroovyJs                                                                          string `json:"codemirror/mode/groovy/groovy.js"`
			CodemirrorModeHamlHamlJs                                                                              string `json:"codemirror/mode/haml/haml.js"`
			CodemirrorModeHandlebarsHandlebarsJs                                                                  string `json:"codemirror/mode/handlebars/handlebars.js"`
			CodemirrorModeHaskellLiterateHaskellLiterateJs                                                        string `json:"codemirror/mode/haskell-literate/haskell-literate.js"`
			CodemirrorModeHaskellHaskellJs                                                                        string `json:"codemirror/mode/haskell/haskell.js"`
			CodemirrorModeHaxeHaxeJs                                                                              string `json:"codemirror/mode/haxe/haxe.js"`
			CodemirrorModeHtmlembeddedHtmlembeddedJs                                                              string `json:"codemirror/mode/htmlembedded/htmlembedded.js"`
			CodemirrorModeHtmlmixedHtmlmixedJs                                                                    string `json:"codemirror/mode/htmlmixed/htmlmixed.js"`
			CodemirrorModeHTTPHTTPJs                                                                              string `json:"codemirror/mode/http/http.js"`
			CodemirrorModeIdlIdlJs                                                                                string `json:"codemirror/mode/idl/idl.js"`
			CodemirrorModeJavascriptJavascriptJs                                                                  string `json:"codemirror/mode/javascript/javascript.js"`
			CodemirrorModeJinja2Jinja2Js                                                                          string `json:"codemirror/mode/jinja2/jinja2.js"`
			CodemirrorModeJsxJsxJs                                                                                string `json:"codemirror/mode/jsx/jsx.js"`
			CodemirrorModeJuliaJuliaJs                                                                            string `json:"codemirror/mode/julia/julia.js"`
			CodemirrorModeLivescriptLivescriptJs                                                                  string `json:"codemirror/mode/livescript/livescript.js"`
			CodemirrorModeLuaLuaJs                                                                                string `json:"codemirror/mode/lua/lua.js"`
			CodemirrorModeMarkdownMarkdownJs                                                                      string `json:"codemirror/mode/markdown/markdown.js"`
			CodemirrorModeMathematicaMathematicaJs                                                                string `json:"codemirror/mode/mathematica/mathematica.js"`
			CodemirrorModeMboxMboxJs                                                                              string `json:"codemirror/mode/mbox/mbox.js"`
			CodemirrorModeMetaJs                                                                                  string `json:"codemirror/mode/meta.js"`
			CodemirrorModeMircMircJs                                                                              string `json:"codemirror/mode/mirc/mirc.js"`
			CodemirrorModeMllikeMllikeJs                                                                          string `json:"codemirror/mode/mllike/mllike.js"`
			CodemirrorModeModelicaModelicaJs                                                                      string `json:"codemirror/mode/modelica/modelica.js"`
			CodemirrorModeMscgenMscgenJs                                                                          string `json:"codemirror/mode/mscgen/mscgen.js"`
			CodemirrorModeMumpsMumpsJs                                                                            string `json:"codemirror/mode/mumps/mumps.js"`
			CodemirrorModeNginxNginxJs                                                                            string `json:"codemirror/mode/nginx/nginx.js"`
			CodemirrorModeNsisNsisJs                                                                              string `json:"codemirror/mode/nsis/nsis.js"`
			CodemirrorModeNtriplesNtriplesJs                                                                      string `json:"codemirror/mode/ntriples/ntriples.js"`
			CodemirrorModeOctaveOctaveJs                                                                          string `json:"codemirror/mode/octave/octave.js"`
			CodemirrorModeOzOzJs                                                                                  string `json:"codemirror/mode/oz/oz.js"`
			CodemirrorModePascalPascalJs                                                                          string `json:"codemirror/mode/pascal/pascal.js"`
			CodemirrorModePegjsPegjsJs                                                                            string `json:"codemirror/mode/pegjs/pegjs.js"`
			CodemirrorModePerlPerlJs                                                                              string `json:"codemirror/mode/perl/perl.js"`
			CodemirrorModePhpPhpJs                                                                                string `json:"codemirror/mode/php/php.js"`
			CodemirrorModePigPigJs                                                                                string `json:"codemirror/mode/pig/pig.js"`
			CodemirrorModePowershellPowershellJs                                                                  string `json:"codemirror/mode/powershell/powershell.js"`
			CodemirrorModePropertiesPropertiesJs                                                                  string `json:"codemirror/mode/properties/properties.js"`
			CodemirrorModeProtobufProtobufJs                                                                      string `json:"codemirror/mode/protobuf/protobuf.js"`
			CodemirrorModePugPugJs                                                                                string `json:"codemirror/mode/pug/pug.js"`
			CodemirrorModePuppetPuppetJs                                                                          string `json:"codemirror/mode/puppet/puppet.js"`
			CodemirrorModePythonPythonJs                                                                          string `json:"codemirror/mode/python/python.js"`
			CodemirrorModeQQJs                                                                                    string `json:"codemirror/mode/q/q.js"`
			CodemirrorModeRRJs                                                                                    string `json:"codemirror/mode/r/r.js"`
			CodemirrorModeRpmChangesIndexHTML                                                                     string `json:"codemirror/mode/rpm/changes/index.html"`
			CodemirrorModeRpmRpmJs                                                                                string `json:"codemirror/mode/rpm/rpm.js"`
			CodemirrorModeRstRstJs                                                                                string `json:"codemirror/mode/rst/rst.js"`
			CodemirrorModeRubyRubyJs                                                                              string `json:"codemirror/mode/ruby/ruby.js"`
			CodemirrorModeRustRustJs                                                                              string `json:"codemirror/mode/rust/rust.js"`
			CodemirrorModeSasSasJs                                                                                string `json:"codemirror/mode/sas/sas.js"`
			CodemirrorModeSassSassJs                                                                              string `json:"codemirror/mode/sass/sass.js"`
			CodemirrorModeSchemeSchemeJs                                                                          string `json:"codemirror/mode/scheme/scheme.js"`
			CodemirrorModeShellShellJs                                                                            string `json:"codemirror/mode/shell/shell.js"`
			CodemirrorModeSieveSieveJs                                                                            string `json:"codemirror/mode/sieve/sieve.js"`
			CodemirrorModeSlimSlimJs                                                                              string `json:"codemirror/mode/slim/slim.js"`
			CodemirrorModeSmalltalkSmalltalkJs                                                                    string `json:"codemirror/mode/smalltalk/smalltalk.js"`
			CodemirrorModeSmartySmartyJs                                                                          string `json:"codemirror/mode/smarty/smarty.js"`
			CodemirrorModeSolrSolrJs                                                                              string `json:"codemirror/mode/solr/solr.js"`
			CodemirrorModeSoySoyJs                                                                                string `json:"codemirror/mode/soy/soy.js"`
			CodemirrorModeSparqlSparqlJs                                                                          string `json:"codemirror/mode/sparql/sparql.js"`
			CodemirrorModeSpreadsheetSpreadsheetJs                                                                string `json:"codemirror/mode/spreadsheet/spreadsheet.js"`
			CodemirrorModeSqlSqlJs                                                                                string `json:"codemirror/mode/sql/sql.js"`
			CodemirrorModeStexStexJs                                                                              string `json:"codemirror/mode/stex/stex.js"`
			CodemirrorModeStylusStylusJs                                                                          string `json:"codemirror/mode/stylus/stylus.js"`
			CodemirrorModeSwiftSwiftJs                                                                            string `json:"codemirror/mode/swift/swift.js"`
			CodemirrorModeTclTclJs                                                                                string `json:"codemirror/mode/tcl/tcl.js"`
			CodemirrorModeTextileTextileJs                                                                        string `json:"codemirror/mode/textile/textile.js"`
			CodemirrorModeTiddlywikiTiddlywikiCSS                                                                 string `json:"codemirror/mode/tiddlywiki/tiddlywiki.css"`
			CodemirrorModeTiddlywikiTiddlywikiJs                                                                  string `json:"codemirror/mode/tiddlywiki/tiddlywiki.js"`
			CodemirrorModeTikiTikiCSS                                                                             string `json:"codemirror/mode/tiki/tiki.css"`
			CodemirrorModeTikiTikiJs                                                                              string `json:"codemirror/mode/tiki/tiki.js"`
			CodemirrorModeTomlTomlJs                                                                              string `json:"codemirror/mode/toml/toml.js"`
			CodemirrorModeTornadoTornadoJs                                                                        string `json:"codemirror/mode/tornado/tornado.js"`
			CodemirrorModeTroffTroffJs                                                                            string `json:"codemirror/mode/troff/troff.js"`
			CodemirrorModeTtcnCfgTtcnCfgJs                                                                        string `json:"codemirror/mode/ttcn-cfg/ttcn-cfg.js"`
			CodemirrorModeTtcnTtcnJs                                                                              string `json:"codemirror/mode/ttcn/ttcn.js"`
			CodemirrorModeTurtleTurtleJs                                                                          string `json:"codemirror/mode/turtle/turtle.js"`
			CodemirrorModeTwigTwigJs                                                                              string `json:"codemirror/mode/twig/twig.js"`
			CodemirrorModeVbVbJs                                                                                  string `json:"codemirror/mode/vb/vb.js"`
			CodemirrorModeVbscriptVbscriptJs                                                                      string `json:"codemirror/mode/vbscript/vbscript.js"`
			CodemirrorModeVelocityVelocityJs                                                                      string `json:"codemirror/mode/velocity/velocity.js"`
			CodemirrorModeVerilogVerilogJs                                                                        string `json:"codemirror/mode/verilog/verilog.js"`
			CodemirrorModeVhdlVhdlJs                                                                              string `json:"codemirror/mode/vhdl/vhdl.js"`
			CodemirrorModeVueVueJs                                                                                string `json:"codemirror/mode/vue/vue.js"`
			CodemirrorModeWebidlWebidlJs                                                                          string `json:"codemirror/mode/webidl/webidl.js"`
			CodemirrorModeXMLXMLJs                                                                                string `json:"codemirror/mode/xml/xml.js"`
			CodemirrorModeXqueryXqueryJs                                                                          string `json:"codemirror/mode/xquery/xquery.js"`
			CodemirrorModeYacasYacasJs                                                                            string `json:"codemirror/mode/yacas/yacas.js"`
			CodemirrorModeYamlFrontmatterYamlFrontmatterJs                                                        string `json:"codemirror/mode/yaml-frontmatter/yaml-frontmatter.js"`
			CodemirrorModeYamlYamlJs                                                                              string `json:"codemirror/mode/yaml/yaml.js"`
			CodemirrorModeZ80Z80Js                                                                                string `json:"codemirror/mode/z80/z80.js"`
			BalloonMinCSS                                                                                         string `json:"balloon.min.css"`
			ReactProductionMinJs                                                                                  string `json:"react.production.min.js"`
			ReactDomProductionMinJs                                                                               string `json:"react-dom.production.min.js"`
			CustomInputMinCSS                                                                                     string `json:"custom-input.min.css"`
			CustomInputMinJs                                                                                      string `json:"custom-input.min.js"`
			TestcasesMinCSS                                                                                       string `json:"testcases.min.css"`
			TestcasesMinJs                                                                                        string `json:"testcases.min.js"`
			BadgesProblemSolvingLevel1Stars1OthersPng                                                             string `json:"badges/problem-solving_level_1_stars_1_others.png"`
			BadgesProblemSolvingLevel2Stars4OthersPng                                                             string `json:"badges/problem-solving_level_2_stars_4_others.png"`
			BadgesProblemSolvingLevel1Stars1LinkedinPng                                                           string `json:"badges/problem-solving_level_1_stars_1_linkedin.png"`
			BadgesProblemSolvingLevel2Stars4LinkedinPng                                                           string `json:"badges/problem-solving_level_2_stars_4_linkedin.png"`
			BadgesProblemSolvingLevel1Stars2OthersPng                                                             string `json:"badges/problem-solving_level_1_stars_2_others.png"`
			BadgesProblemSolvingLevel2Stars3OthersPng                                                             string `json:"badges/problem-solving_level_2_stars_3_others.png"`
			BadgesProblemSolvingLevel1Stars2LinkedinPng                                                           string `json:"badges/problem-solving_level_1_stars_2_linkedin.png"`
			BadgesProblemSolvingLevel2Stars3LinkedinPng                                                           string `json:"badges/problem-solving_level_2_stars_3_linkedin.png"`
			BadgesProblemSolvingLevel3Stars5OthersPng                                                             string `json:"badges/problem-solving_level_3_stars_5_others.png"`
			BadgesProblemSolvingLevel3Stars6OthersPng                                                             string `json:"badges/problem-solving_level_3_stars_6_others.png"`
			BadgesProblemSolvingLevel3Stars5LinkedinPng                                                           string `json:"badges/problem-solving_level_3_stars_5_linkedin.png"`
			BadgesCppLevel1Stars1OthersPng                                                                        string `json:"badges/cpp_level_1_stars_1_others.png"`
			BadgesProblemSolvingLevel3Stars6LinkedinPng                                                           string `json:"badges/problem-solving_level_3_stars_6_linkedin.png"`
			BadgesCppLevel1Stars2OthersPng                                                                        string `json:"badges/cpp_level_1_stars_2_others.png"`
			BadgesCppLevel1Stars1LinkedinPng                                                                      string `json:"badges/cpp_level_1_stars_1_linkedin.png"`
			BadgesCppLevel1Stars2LinkedinPng                                                                      string `json:"badges/cpp_level_1_stars_2_linkedin.png"`
			BadgesCppLevel2Stars3OthersPng                                                                        string `json:"badges/cpp_level_2_stars_3_others.png"`
			BadgesCppLevel2Stars4OthersPng                                                                        string `json:"badges/cpp_level_2_stars_4_others.png"`
			BadgesCppLevel2Stars3LinkedinPng                                                                      string `json:"badges/cpp_level_2_stars_3_linkedin.png"`
			BadgesCppLevel3Stars5OthersPng                                                                        string `json:"badges/cpp_level_3_stars_5_others.png"`
			BadgesCppLevel2Stars4LinkedinPng                                                                      string `json:"badges/cpp_level_2_stars_4_linkedin.png"`
			BadgesPythonLevel1Stars1OthersPng                                                                     string `json:"badges/python_level_1_stars_1_others.png"`
			BadgesCppLevel3Stars5LinkedinPng                                                                      string `json:"badges/cpp_level_3_stars_5_linkedin.png"`
			BadgesPythonLevel1Stars2OthersPng                                                                     string `json:"badges/python_level_1_stars_2_others.png"`
			BadgesPythonLevel1Stars1LinkedinPng                                                                   string `json:"badges/python_level_1_stars_1_linkedin.png"`
			BadgesPythonLevel2Stars3OthersPng                                                                     string `json:"badges/python_level_2_stars_3_others.png"`
			BadgesPythonLevel1Stars2LinkedinPng                                                                   string `json:"badges/python_level_1_stars_2_linkedin.png"`
			BadgesPythonLevel2Stars4OthersPng                                                                     string `json:"badges/python_level_2_stars_4_others.png"`
			BadgesPythonLevel2Stars3LinkedinPng                                                                   string `json:"badges/python_level_2_stars_3_linkedin.png"`
			BadgesPythonLevel3Stars5OthersPng                                                                     string `json:"badges/python_level_3_stars_5_others.png"`
			BadgesPythonLevel2Stars4LinkedinPng                                                                   string `json:"badges/python_level_2_stars_4_linkedin.png"`
			BadgesRubyLevel1Stars1OthersPng                                                                       string `json:"badges/ruby_level_1_stars_1_others.png"`
			BadgesRubyLevel1Stars2OthersPng                                                                       string `json:"badges/ruby_level_1_stars_2_others.png"`
			BadgesPythonLevel3Stars5LinkedinPng                                                                   string `json:"badges/python_level_3_stars_5_linkedin.png"`
			BadgesRubyLevel1Stars1LinkedinPng                                                                     string `json:"badges/ruby_level_1_stars_1_linkedin.png"`
			BadgesRubyLevel2Stars3OthersPng                                                                       string `json:"badges/ruby_level_2_stars_3_others.png"`
			BadgesRubyLevel1Stars2LinkedinPng                                                                     string `json:"badges/ruby_level_1_stars_2_linkedin.png"`
			BadgesRubyLevel2Stars4OthersPng                                                                       string `json:"badges/ruby_level_2_stars_4_others.png"`
			BadgesRubyLevel2Stars3LinkedinPng                                                                     string `json:"badges/ruby_level_2_stars_3_linkedin.png"`
			BadgesRubyLevel3Stars5OthersPng                                                                       string `json:"badges/ruby_level_3_stars_5_others.png"`
			BadgesRubyLevel2Stars4LinkedinPng                                                                     string `json:"badges/ruby_level_2_stars_4_linkedin.png"`
			BadgesSqlLevel1Stars1OthersPng                                                                        string `json:"badges/sql_level_1_stars_1_others.png"`
			BadgesRubyLevel3Stars5LinkedinPng                                                                     string `json:"badges/ruby_level_3_stars_5_linkedin.png"`
			BadgesSqlLevel1Stars2OthersPng                                                                        string `json:"badges/sql_level_1_stars_2_others.png"`
			BadgesSqlLevel1Stars1LinkedinPng                                                                      string `json:"badges/sql_level_1_stars_1_linkedin.png"`
			BadgesSqlLevel2Stars3OthersPng                                                                        string `json:"badges/sql_level_2_stars_3_others.png"`
			BadgesSqlLevel1Stars2LinkedinPng                                                                      string `json:"badges/sql_level_1_stars_2_linkedin.png"`
			BadgesSqlLevel2Stars4OthersPng                                                                        string `json:"badges/sql_level_2_stars_4_others.png"`
			BadgesSqlLevel2Stars3LinkedinPng                                                                      string `json:"badges/sql_level_2_stars_3_linkedin.png"`
			BadgesSqlLevel3Stars5OthersPng                                                                        string `json:"badges/sql_level_3_stars_5_others.png"`
			BadgesSqlLevel2Stars4LinkedinPng                                                                      string `json:"badges/sql_level_2_stars_4_linkedin.png"`
			BadgesJavaLevel1Stars1OthersPng                                                                       string `json:"badges/java_level_1_stars_1_others.png"`
			BadgesSqlLevel3Stars5LinkedinPng                                                                      string `json:"badges/sql_level_3_stars_5_linkedin.png"`
			BadgesJavaLevel1Stars2OthersPng                                                                       string `json:"badges/java_level_1_stars_2_others.png"`
			BadgesJavaLevel1Stars1LinkedinPng                                                                     string `json:"badges/java_level_1_stars_1_linkedin.png"`
			BadgesJavaLevel2Stars3OthersPng                                                                       string `json:"badges/java_level_2_stars_3_others.png"`
			BadgesJavaLevel1Stars2LinkedinPng                                                                     string `json:"badges/java_level_1_stars_2_linkedin.png"`
			BadgesJavaLevel2Stars4OthersPng                                                                       string `json:"badges/java_level_2_stars_4_others.png"`
			BadgesJavaLevel2Stars3LinkedinPng                                                                     string `json:"badges/java_level_2_stars_3_linkedin.png"`
			BadgesJavaLevel3Stars5OthersPng                                                                       string `json:"badges/java_level_3_stars_5_others.png"`
			BadgesJavaLevel2Stars4LinkedinPng                                                                     string `json:"badges/java_level_2_stars_4_linkedin.png"`
			Badges30DaysOfCodeLevel1Stars1OthersPng                                                               string `json:"badges/30-days-of-code_level_1_stars_1_others.png"`
			BadgesJavaLevel3Stars5LinkedinPng                                                                     string `json:"badges/java_level_3_stars_5_linkedin.png"`
			Badges30DaysOfCodeLevel2Stars3OthersPng                                                               string `json:"badges/30-days-of-code_level_2_stars_3_others.png"`
			Badges30DaysOfCodeLevel1Stars2OthersPng                                                               string `json:"badges/30-days-of-code_level_1_stars_2_others.png"`
			Badges30DaysOfCodeLevel1Stars1LinkedinPng                                                             string `json:"badges/30-days-of-code_level_1_stars_1_linkedin.png"`
			Badges30DaysOfCodeLevel1Stars2LinkedinPng                                                             string `json:"badges/30-days-of-code_level_1_stars_2_linkedin.png"`
			Badges30DaysOfCodeLevel2Stars4OthersPng                                                               string `json:"badges/30-days-of-code_level_2_stars_4_others.png"`
			Badges30DaysOfCodeLevel2Stars3LinkedinPng                                                             string `json:"badges/30-days-of-code_level_2_stars_3_linkedin.png"`
			Badges30DaysOfCodeLevel3Stars5OthersPng                                                               string `json:"badges/30-days-of-code_level_3_stars_5_others.png"`
			Badges30DaysOfCodeLevel2Stars4LinkedinPng                                                             string `json:"badges/30-days-of-code_level_2_stars_4_linkedin.png"`
			Badges10DaysOfJavascriptLevel1Stars1OthersPng                                                         string `json:"badges/10-days-of-javascript_level_1_stars_1_others.png"`
			Badges30DaysOfCodeLevel3Stars5LinkedinPng                                                             string `json:"badges/30-days-of-code_level_3_stars_5_linkedin.png"`
			Badges10DaysOfJavascriptLevel1Stars2OthersPng                                                         string `json:"badges/10-days-of-javascript_level_1_stars_2_others.png"`
			Badges10DaysOfJavascriptLevel1Stars1LinkedinPng                                                       string `json:"badges/10-days-of-javascript_level_1_stars_1_linkedin.png"`
			Badges10DaysOfJavascriptLevel1Stars2LinkedinPng                                                       string `json:"badges/10-days-of-javascript_level_1_stars_2_linkedin.png"`
			Badges10DaysOfJavascriptLevel2Stars3OthersPng                                                         string `json:"badges/10-days-of-javascript_level_2_stars_3_others.png"`
			Badges10DaysOfJavascriptLevel2Stars4OthersPng                                                         string `json:"badges/10-days-of-javascript_level_2_stars_4_others.png"`
			Badges10DaysOfJavascriptLevel2Stars3LinkedinPng                                                       string `json:"badges/10-days-of-javascript_level_2_stars_3_linkedin.png"`
			Badges10DaysOfJavascriptLevel3Stars5OthersPng                                                         string `json:"badges/10-days-of-javascript_level_3_stars_5_others.png"`
			Badges10DaysOfJavascriptLevel2Stars4LinkedinPng                                                       string `json:"badges/10-days-of-javascript_level_2_stars_4_linkedin.png"`
			Badges10DaysOfStatisticsLevel1Stars1OthersPng                                                         string `json:"badges/10-days-of-statistics_level_1_stars_1_others.png"`
			Badges10DaysOfJavascriptLevel3Stars5LinkedinPng                                                       string `json:"badges/10-days-of-javascript_level_3_stars_5_linkedin.png"`
			Badges10DaysOfStatisticsLevel1Stars2OthersPng                                                         string `json:"badges/10-days-of-statistics_level_1_stars_2_others.png"`
			Badges10DaysOfStatisticsLevel1Stars1LinkedinPng                                                       string `json:"badges/10-days-of-statistics_level_1_stars_1_linkedin.png"`
			Badges10DaysOfStatisticsLevel2Stars3OthersPng                                                         string `json:"badges/10-days-of-statistics_level_2_stars_3_others.png"`
			Badges10DaysOfStatisticsLevel1Stars2LinkedinPng                                                       string `json:"badges/10-days-of-statistics_level_1_stars_2_linkedin.png"`
			Badges10DaysOfStatisticsLevel2Stars4OthersPng                                                         string `json:"badges/10-days-of-statistics_level_2_stars_4_others.png"`
			Badges10DaysOfStatisticsLevel2Stars3LinkedinPng                                                       string `json:"badges/10-days-of-statistics_level_2_stars_3_linkedin.png"`
			Badges10DaysOfStatisticsLevel3Stars5OthersPng                                                         string `json:"badges/10-days-of-statistics_level_3_stars_5_others.png"`
			CodemirrorBasicJs                                                                                     string `json:"codemirror_basic.js"`
			PolyfillMinJs                                                                                         string `json:"polyfill.min.js"`
			HackerrankRVendorJs                                                                                   string `json:"hackerrank_r_vendor.js"`
			VendorsHackerrankRChallengeListHackerrankRChallengeListV2HackerrankRInterviewHackerrank1976138AJs     string `json:"vendors~hackerrank_r_challenge_list~hackerrank_r_challenge_list_v2~hackerrank_r_interview~hackerrank~1976138a.js"`
			VendorsHackerrankRChallengeListHackerrankRChallengeListV2HackerrankRInterviewHackerrank1976138AJsMap  string `json:"vendors~hackerrank_r_challenge_list~hackerrank_r_challenge_list_v2~hackerrank_r_interview~hackerrank~1976138a.js.map"`
			RuntimeJs                                                                                             string `json:"runtime.js"`
			RuntimeJsMap                                                                                          string `json:"runtime.js.map"`
			VendorsCustominputHackerrankRChallengeHackerrankRCodeSnippetsHackerrankROnboardingCSS                 string `json:"vendors~custominput~hackerrank_r_challenge~hackerrank_r_code_snippets~hackerrank_r_onboarding.css"`
			VendorsCustominputHackerrankRChallengeHackerrankRCodeSnippetsHackerrankROnboardingJs                  string `json:"vendors~custominput~hackerrank_r_challenge~hackerrank_r_code_snippets~hackerrank_r_onboarding.js"`
			VendorsCustominputHackerrankRChallengeHackerrankRCodeSnippetsHackerrankROnboardingCSSMap              string `json:"vendors~custominput~hackerrank_r_challenge~hackerrank_r_code_snippets~hackerrank_r_onboarding.css.map"`
			VendorsCustominputHackerrankRChallengeHackerrankRCodeSnippetsHackerrankROnboardingJsMap               string `json:"vendors~custominput~hackerrank_r_challenge~hackerrank_r_code_snippets~hackerrank_r_onboarding.js.map"`
			VendorsHackerrankRChallengeListHackerrankRChallengeListV2HackerrankRInterviewHackerrank14Aa502AJs     string `json:"vendors~hackerrank_r_challenge_list~hackerrank_r_challenge_list_v2~hackerrank_r_interview~hackerrank~14aa502a.js"`
			VendorsHackerrankRChallengeListHackerrankRChallengeListV2HackerrankRInterviewHackerrank14Aa502AJsMap  string `json:"vendors~hackerrank_r_challenge_list~hackerrank_r_challenge_list_v2~hackerrank_r_interview~hackerrank~14aa502a.js.map"`
			VendorsHackerrankRCommunityHackerrankRSourcingApplicationsHackerrankRSourcingCompanyJobs9Ddfcdb2Js    string `json:"vendors~hackerrank_r_community~hackerrank_r_sourcing_applications~hackerrank_r_sourcing_company_jobs~9ddfcdb2.js"`
			VendorsHackerrankRCommunityHackerrankRSourcingApplicationsHackerrankRSourcingCompanyJobs9Ddfcdb2JsMap string `json:"vendors~hackerrank_r_community~hackerrank_r_sourcing_applications~hackerrank_r_sourcing_company_jobs~9ddfcdb2.js.map"`
			Five2Da6B68Bc6E31Ad3E665Js                                                                            string `json:"5-2da6b68bc6e31ad3e665.js"`
			Sourcemaps52Da6B68Bc6E31Ad3E665JsMap                                                                  string `json:"sourcemaps/5-2da6b68bc6e31ad3e665.js.map"`
			HackerrankRChallengeListHackerrankRChallengeListV2HackerrankRInterviewCSS                             string `json:"hackerrank_r_challenge_list~hackerrank_r_challenge_list_v2~hackerrank_r_interview.css"`
			HackerrankRChallengeListHackerrankRChallengeListV2HackerrankRInterviewJs                              string `json:"hackerrank_r_challenge_list~hackerrank_r_challenge_list_v2~hackerrank_r_interview.js"`
			HackerrankRChallengeListHackerrankRChallengeListV2HackerrankRInterviewCSSMap                          string `json:"hackerrank_r_challenge_list~hackerrank_r_challenge_list_v2~hackerrank_r_interview.css.map"`
			HackerrankRChallengeListHackerrankRChallengeListV2HackerrankRInterviewJsMap                           string `json:"hackerrank_r_challenge_list~hackerrank_r_challenge_list_v2~hackerrank_r_interview.js.map"`
			VendorsHackerrankRChallengeHackerrankRCodeSnippetsHackerrankROnboardingCSS                            string `json:"vendors~hackerrank_r_challenge~hackerrank_r_code_snippets~hackerrank_r_onboarding.css"`
			VendorsHackerrankRChallengeHackerrankRCodeSnippetsHackerrankROnboardingJs                             string `json:"vendors~hackerrank_r_challenge~hackerrank_r_code_snippets~hackerrank_r_onboarding.js"`
			VendorsHackerrankRChallengeHackerrankRCodeSnippetsHackerrankROnboardingCSSMap                         string `json:"vendors~hackerrank_r_challenge~hackerrank_r_code_snippets~hackerrank_r_onboarding.css.map"`
			VendorsHackerrankRChallengeHackerrankRCodeSnippetsHackerrankROnboardingJsMap                          string `json:"vendors~hackerrank_r_challenge~hackerrank_r_code_snippets~hackerrank_r_onboarding.js.map"`
			VendorsHackerrankRCommunityHackerrankRSourcingApplicationsHackerrankRSourcingCompanyJobsCSS           string `json:"vendors~hackerrank_r_community~hackerrank_r_sourcing_applications~hackerrank_r_sourcing_company_jobs.css"`
			VendorsHackerrankRCommunityHackerrankRSourcingApplicationsHackerrankRSourcingCompanyJobsJs            string `json:"vendors~hackerrank_r_community~hackerrank_r_sourcing_applications~hackerrank_r_sourcing_company_jobs.js"`
			VendorsHackerrankRCommunityHackerrankRSourcingApplicationsHackerrankRSourcingCompanyJobsCSSMap        string `json:"vendors~hackerrank_r_community~hackerrank_r_sourcing_applications~hackerrank_r_sourcing_company_jobs.css.map"`
			VendorsHackerrankRCommunityHackerrankRSourcingApplicationsHackerrankRSourcingCompanyJobsJsMap         string `json:"vendors~hackerrank_r_community~hackerrank_r_sourcing_applications~hackerrank_r_sourcing_company_jobs.js.map"`
			VendorsHackerrankRCommunityHackerrankRWorkLspClientJs                                                 string `json:"vendors~hackerrank_r_community~hackerrank_r_work~lsp_client.js"`
			VendorsHackerrankRCommunityHackerrankRWorkLspClientJsMap                                              string `json:"vendors~hackerrank_r_community~hackerrank_r_work~lsp_client.js.map"`
			VendorsHackerrankRJobsHackerrankRTestinviteHackerrankRTestsettingsJs                                  string `json:"vendors~hackerrank_r_jobs~hackerrank_r_testinvite~hackerrank_r_testsettings.js"`
			VendorsHackerrankRJobsHackerrankRTestinviteHackerrankRTestsettingsJsMap                               string `json:"vendors~hackerrank_r_jobs~hackerrank_r_testinvite~hackerrank_r_testsettings.js.map"`
			HackerrankRCodeSnippetsHackerrankROnboardingCSS                                                       string `json:"hackerrank_r_code_snippets~hackerrank_r_onboarding.css"`
			HackerrankRCodeSnippetsHackerrankROnboardingJs                                                        string `json:"hackerrank_r_code_snippets~hackerrank_r_onboarding.js"`
			HackerrankRCodeSnippetsHackerrankROnboardingCSSMap                                                    string `json:"hackerrank_r_code_snippets~hackerrank_r_onboarding.css.map"`
			HackerrankRCodeSnippetsHackerrankROnboardingJsMap                                                     string `json:"hackerrank_r_code_snippets~hackerrank_r_onboarding.js.map"`
			HackerrankRLeaderboardHackerrankRLeaderboardv2CSS                                                     string `json:"hackerrank_r_leaderboard~hackerrank_r_leaderboardv2.css"`
			HackerrankRLeaderboardHackerrankRLeaderboardv2Js                                                      string `json:"hackerrank_r_leaderboard~hackerrank_r_leaderboardv2.js"`
			HackerrankRLeaderboardHackerrankRLeaderboardv2CSSMap                                                  string `json:"hackerrank_r_leaderboard~hackerrank_r_leaderboardv2.css.map"`
			HackerrankRLeaderboardHackerrankRLeaderboardv2JsMap                                                   string `json:"hackerrank_r_leaderboard~hackerrank_r_leaderboardv2.js.map"`
			HackerrankRLibraryHackerrankRTestquestionsCSS                                                         string `json:"hackerrank_r_library~hackerrank_r_testquestions.css"`
			HackerrankRLibraryHackerrankRTestquestionsJs                                                          string `json:"hackerrank_r_library~hackerrank_r_testquestions.js"`
			HackerrankRLibraryHackerrankRTestquestionsCSSMap                                                      string `json:"hackerrank_r_library~hackerrank_r_testquestions.css.map"`
			HackerrankRLibraryHackerrankRTestquestionsJsMap                                                       string `json:"hackerrank_r_library~hackerrank_r_testquestions.js.map"`
			HackerrankRProfileV2HackerrankRSourcingCompanyCSS                                                     string `json:"hackerrank_r_profile_v2~hackerrank_r_sourcing_company.css"`
			HackerrankRProfileV2HackerrankRSourcingCompanyJs                                                      string `json:"hackerrank_r_profile_v2~hackerrank_r_sourcing_company.js"`
			HackerrankRProfileV2HackerrankRSourcingCompanyCSSMap                                                  string `json:"hackerrank_r_profile_v2~hackerrank_r_sourcing_company.css.map"`
			HackerrankRProfileV2HackerrankRSourcingCompanyJsMap                                                   string `json:"hackerrank_r_profile_v2~hackerrank_r_sourcing_company.js.map"`
			VendorsHackerrankRChallengeHackerrankROnboardingJs                                                    string `json:"vendors~hackerrank_r_challenge~hackerrank_r_onboarding.js"`
			VendorsHackerrankRChallengeHackerrankROnboardingJsMap                                                 string `json:"vendors~hackerrank_r_challenge~hackerrank_r_onboarding.js.map"`
			VendorsHackerrankRChallengeHackerrankRWorkCSS                                                         string `json:"vendors~hackerrank_r_challenge~hackerrank_r_work.css"`
			VendorsHackerrankRChallengeHackerrankRWorkJs                                                          string `json:"vendors~hackerrank_r_challenge~hackerrank_r_work.js"`
			VendorsHackerrankRChallengeHackerrankRWorkCSSMap                                                      string `json:"vendors~hackerrank_r_challenge~hackerrank_r_work.css.map"`
			VendorsHackerrankRChallengeHackerrankRWorkJsMap                                                       string `json:"vendors~hackerrank_r_challenge~hackerrank_r_work.js.map"`
			VendorsHackerrankRCodeSnippetsHackerrankRForumJs                                                      string `json:"vendors~hackerrank_r_code_snippets~hackerrank_r_forum.js"`
			VendorsHackerrankRCodeSnippetsHackerrankRForumJsMap                                                   string `json:"vendors~hackerrank_r_code_snippets~hackerrank_r_forum.js.map"`
			VendorsHackerrankRCodeshellLibHackerrankRTeamsJs                                                      string `json:"vendors~hackerrank_r_codeshell_lib~hackerrank_r_teams.js"`
			VendorsHackerrankRCodeshellLibHackerrankRTeamsJsMap                                                   string `json:"vendors~hackerrank_r_codeshell_lib~hackerrank_r_teams.js.map"`
			VendorsHackerrankRCommunityHackerrankRWorkCSS                                                         string `json:"vendors~hackerrank_r_community~hackerrank_r_work.css"`
			VendorsHackerrankRCommunityHackerrankRWorkJs                                                          string `json:"vendors~hackerrank_r_community~hackerrank_r_work.js"`
			VendorsHackerrankRCommunityHackerrankRWorkCSSMap                                                      string `json:"vendors~hackerrank_r_community~hackerrank_r_work.css.map"`
			VendorsHackerrankRCommunityHackerrankRWorkJsMap                                                       string `json:"vendors~hackerrank_r_community~hackerrank_r_work.js.map"`
			VendorsHackerrankRDashboardHackerrankRWorkJs                                                          string `json:"vendors~hackerrank_r_dashboard~hackerrank_r_work.js"`
			VendorsHackerrankRDashboardHackerrankRWorkJsMap                                                       string `json:"vendors~hackerrank_r_dashboard~hackerrank_r_work.js.map"`
			VendorsHackerrankRProfileV2HackerrankRSourcingCompanyCSS                                              string `json:"vendors~hackerrank_r_profile_v2~hackerrank_r_sourcing_company.css"`
			VendorsHackerrankRProfileV2HackerrankRSourcingCompanyJs                                               string `json:"vendors~hackerrank_r_profile_v2~hackerrank_r_sourcing_company.js"`
			VendorsHackerrankRProfileV2HackerrankRSourcingCompanyCSSMap                                           string `json:"vendors~hackerrank_r_profile_v2~hackerrank_r_sourcing_company.css.map"`
			VendorsHackerrankRProfileV2HackerrankRSourcingCompanyJsMap                                            string `json:"vendors~hackerrank_r_profile_v2~hackerrank_r_sourcing_company.js.map"`
			VendorsHackerrankRTestinviteHackerrankRTestsettingsJs                                                 string `json:"vendors~hackerrank_r_testinvite~hackerrank_r_testsettings.js"`
			VendorsHackerrankRTestinviteHackerrankRTestsettingsJsMap                                              string `json:"vendors~hackerrank_r_testinvite~hackerrank_r_testsettings.js.map"`
			VendorsHackerrankRTestquestionsHackerrankRTestsettingsJs                                              string `json:"vendors~hackerrank_r_testquestions~hackerrank_r_testsettings.js"`
			VendorsHackerrankRTestquestionsHackerrankRTestsettingsJsMap                                           string `json:"vendors~hackerrank_r_testquestions~hackerrank_r_testsettings.js.map"`
			Two4337362Eacf0A3Abf26EdJs                                                                            string `json:"24-337362eacf0a3abf26ed.js"`
			Sourcemaps24337362Eacf0A3Abf26EdJsMap                                                                 string `json:"sourcemaps/24-337362eacf0a3abf26ed.js.map"`
			Two526065B48D9F14Cd0De3AJs                                                                            string `json:"25-26065b48d9f14cd0de3a.js"`
			Sourcemaps2526065B48D9F14Cd0De3AJsMap                                                                 string `json:"sourcemaps/25-26065b48d9f14cd0de3a.js.map"`
			Two682F3Ad8280194A812230Js                                                                            string `json:"26-82f3ad8280194a812230.js"`
			Sourcemaps2682F3Ad8280194A812230JsMap                                                                 string `json:"sourcemaps/26-82f3ad8280194a812230.js.map"`
			Two7D2D769712Ec60B1AedddJs                                                                            string `json:"27-d2d769712ec60b1aeddd.js"`
			Sourcemaps27D2D769712Ec60B1AedddJsMap                                                                 string `json:"sourcemaps/27-d2d769712ec60b1aeddd.js.map"`
			Two8E404D63B30A326Fee91BJs                                                                            string `json:"28-e404d63b30a326fee91b.js"`
			Sourcemaps28E404D63B30A326Fee91BJsMap                                                                 string `json:"sourcemaps/28-e404d63b30a326fee91b.js.map"`
			Two906D1581F4E23667Ff434Js                                                                            string `json:"29-06d1581f4e23667ff434.js"`
			Sourcemaps2906D1581F4E23667Ff434JsMap                                                                 string `json:"sourcemaps/29-06d1581f4e23667ff434.js.map"`
			Three0Ea4B6Da7450D3989F1E0Js                                                                          string `json:"30-ea4b6da7450d3989f1e0.js"`
			Sourcemaps30Ea4B6Da7450D3989F1E0JsMap                                                                 string `json:"sourcemaps/30-ea4b6da7450d3989f1e0.js.map"`
			Three1064Fbfe640C12E8E5Dd0Js                                                                          string `json:"31-064fbfe640c12e8e5dd0.js"`
			Sourcemaps31064Fbfe640C12E8E5Dd0JsMap                                                                 string `json:"sourcemaps/31-064fbfe640c12e8e5dd0.js.map"`
			Three205F5Dbe3C4Fbf1C70571Js                                                                          string `json:"32-05f5dbe3c4fbf1c70571.js"`
			Sourcemaps3205F5Dbe3C4Fbf1C70571JsMap                                                                 string `json:"sourcemaps/32-05f5dbe3c4fbf1c70571.js.map"`
			Three3258A6763A9174Fa2D889Js                                                                          string `json:"33-258a6763a9174fa2d889.js"`
			Sourcemaps33258A6763A9174Fa2D889JsMap                                                                 string `json:"sourcemaps/33-258a6763a9174fa2d889.js.map"`
			Three4D613090Ba5Af8Db2D87EJs                                                                          string `json:"34-d613090ba5af8db2d87e.js"`
			Sourcemaps34D613090Ba5Af8Db2D87EJsMap                                                                 string `json:"sourcemaps/34-d613090ba5af8db2d87e.js.map"`
			Three58D2D5E68Dcd7Bf7F3Bb9Js                                                                          string `json:"35-8d2d5e68dcd7bf7f3bb9.js"`
			Sourcemaps358D2D5E68Dcd7Bf7F3Bb9JsMap                                                                 string `json:"sourcemaps/35-8d2d5e68dcd7bf7f3bb9.js.map"`
			Three6Ea2929Da821Ece8740EbJs                                                                          string `json:"36-ea2929da821ece8740eb.js"`
			Sourcemaps36Ea2929Da821Ece8740EbJsMap                                                                 string `json:"sourcemaps/36-ea2929da821ece8740eb.js.map"`
			Three717C786536A3850C47FebJs                                                                          string `json:"37-17c786536a3850c47feb.js"`
			Sourcemaps3717C786536A3850C47FebJsMap                                                                 string `json:"sourcemaps/37-17c786536a3850c47feb.js.map"`
			Three83E3709C298D880279BbcJs                                                                          string `json:"38-3e3709c298d880279bbc.js"`
			Sourcemaps383E3709C298D880279BbcJsMap                                                                 string `json:"sourcemaps/38-3e3709c298d880279bbc.js.map"`
			Three9C210Eec841E7105F337BJs                                                                          string `json:"39-c210eec841e7105f337b.js"`
			Sourcemaps39C210Eec841E7105F337BJsMap                                                                 string `json:"sourcemaps/39-c210eec841e7105f337b.js.map"`
			Four02B5C07Cac7A62Ed29C5DJs                                                                           string `json:"40-2b5c07cac7a62ed29c5d.js"`
			Sourcemaps402B5C07Cac7A62Ed29C5DJsMap                                                                 string `json:"sourcemaps/40-2b5c07cac7a62ed29c5d.js.map"`
			Four16Beb3D869C0D9A03280EJs                                                                           string `json:"41-6beb3d869c0d9a03280e.js"`
			Sourcemaps416Beb3D869C0D9A03280EJsMap                                                                 string `json:"sourcemaps/41-6beb3d869c0d9a03280e.js.map"`
			Four25Be4Ea0282Bc4D734E64Js                                                                           string `json:"42-5be4ea0282bc4d734e64.js"`
			Sourcemaps425Be4Ea0282Bc4D734E64JsMap                                                                 string `json:"sourcemaps/42-5be4ea0282bc4d734e64.js.map"`
			Four36917F523E9Dcc3Cbd494Js                                                                           string `json:"43-6917f523e9dcc3cbd494.js"`
			Sourcemaps436917F523E9Dcc3Cbd494JsMap                                                                 string `json:"sourcemaps/43-6917f523e9dcc3cbd494.js.map"`
			Four4Fe8E4208873B0Ea82026Js                                                                           string `json:"44-fe8e4208873b0ea82026.js"`
			Sourcemaps44Fe8E4208873B0Ea82026JsMap                                                                 string `json:"sourcemaps/44-fe8e4208873b0ea82026.js.map"`
			Four59Dd7Db261075F996197BJs                                                                           string `json:"45-9dd7db261075f996197b.js"`
			Sourcemaps459Dd7Db261075F996197BJsMap                                                                 string `json:"sourcemaps/45-9dd7db261075f996197b.js.map"`
			CodeshellEditorThemeMCSS                                                                              string `json:"codeshell_editor_theme_m.css"`
			CodeshellEditorThemeMJs                                                                               string `json:"codeshell_editor_theme_m.js"`
			CodeshellEditorThemeMCSSMap                                                                           string `json:"codeshell_editor_theme_m.css.map"`
			CodeshellEditorThemeMJsMap                                                                            string `json:"codeshell_editor_theme_m.js.map"`
			CustominputJs                                                                                         string `json:"custominput.js"`
			CustominputJsMap                                                                                      string `json:"custominput.js.map"`
			FirebaseClientJs                                                                                      string `json:"firebase_client.js"`
			FirebaseClientJsMap                                                                                   string `json:"firebase_client.js.map"`
			HackerrankRAppCSS                                                                                     string `json:"hackerrank_r_app.css"`
			HackerrankRAppJs                                                                                      string `json:"hackerrank_r_app.js"`
			HackerrankRAppCSSMap                                                                                  string `json:"hackerrank_r_app.css.map"`
			HackerrankRAppJsMap                                                                                   string `json:"hackerrank_r_app.js.map"`
			HackerrankRAuthCSS                                                                                    string `json:"hackerrank_r_auth.css"`
			HackerrankRAuthJs                                                                                     string `json:"hackerrank_r_auth.js"`
			HackerrankRAuthCSSMap                                                                                 string `json:"hackerrank_r_auth.css.map"`
			HackerrankRAuthJsMap                                                                                  string `json:"hackerrank_r_auth.js.map"`
			HackerrankRCalendarCSS                                                                                string `json:"hackerrank_r_calendar.css"`
			HackerrankRCalendarJs                                                                                 string `json:"hackerrank_r_calendar.js"`
			HackerrankRCalendarCSSMap                                                                             string `json:"hackerrank_r_calendar.css.map"`
			HackerrankRCalendarJsMap                                                                              string `json:"hackerrank_r_calendar.js.map"`
			HackerrankRChallengeCSS                                                                               string `json:"hackerrank_r_challenge.css"`
			HackerrankRChallengeJs                                                                                string `json:"hackerrank_r_challenge.js"`
			HackerrankRChallengeCSSMap                                                                            string `json:"hackerrank_r_challenge.css.map"`
			HackerrankRChallengeJsMap                                                                             string `json:"hackerrank_r_challenge.js.map"`
			HackerrankRChallengeListCSS                                                                           string `json:"hackerrank_r_challenge_list.css"`
			HackerrankRChallengeListJs                                                                            string `json:"hackerrank_r_challenge_list.js"`
			HackerrankRChallengeListCSSMap                                                                        string `json:"hackerrank_r_challenge_list.css.map"`
			HackerrankRChallengeListJsMap                                                                         string `json:"hackerrank_r_challenge_list.js.map"`
			HackerrankRChallengeListV2CSS                                                                         string `json:"hackerrank_r_challenge_list_v2.css"`
			HackerrankRChallengeListV2Js                                                                          string `json:"hackerrank_r_challenge_list_v2.js"`
			HackerrankRChallengeListV2CSSMap                                                                      string `json:"hackerrank_r_challenge_list_v2.css.map"`
			HackerrankRChallengeListV2JsMap                                                                       string `json:"hackerrank_r_challenge_list_v2.js.map"`
			HackerrankRCodeSnippetsCSS                                                                            string `json:"hackerrank_r_code_snippets.css"`
			HackerrankRCodeSnippetsJs                                                                             string `json:"hackerrank_r_code_snippets.js"`
			HackerrankRCodeSnippetsCSSMap                                                                         string `json:"hackerrank_r_code_snippets.css.map"`
			HackerrankRCodeSnippetsJsMap                                                                          string `json:"hackerrank_r_code_snippets.js.map"`
			HackerrankRCodeshellLibJs                                                                             string `json:"hackerrank_r_codeshell_lib.js"`
			HackerrankRCodeshellLibJsMap                                                                          string `json:"hackerrank_r_codeshell_lib.js.map"`
			HackerrankRCommunityCSS                                                                               string `json:"hackerrank_r_community.css"`
			HackerrankRCommunityJs                                                                                string `json:"hackerrank_r_community.js"`
			HackerrankRCommunityCSSMap                                                                            string `json:"hackerrank_r_community.css.map"`
			HackerrankRCommunityJsMap                                                                             string `json:"hackerrank_r_community.js.map"`
			HackerrankRCompanyHeaderCSS                                                                           string `json:"hackerrank_r_company_header.css"`
			HackerrankRCompanyHeaderJs                                                                            string `json:"hackerrank_r_company_header.js"`
			HackerrankRCompanyHeaderCSSMap                                                                        string `json:"hackerrank_r_company_header.css.map"`
			HackerrankRCompanyHeaderJsMap                                                                         string `json:"hackerrank_r_company_header.js.map"`
			HackerrankRContestCSS                                                                                 string `json:"hackerrank_r_contest.css"`
			HackerrankRContestJs                                                                                  string `json:"hackerrank_r_contest.js"`
			HackerrankRContestCSSMap                                                                              string `json:"hackerrank_r_contest.css.map"`
			HackerrankRContestJsMap                                                                               string `json:"hackerrank_r_contest.js.map"`
			HackerrankRDashboardCSS                                                                               string `json:"hackerrank_r_dashboard.css"`
			HackerrankRDashboardJs                                                                                string `json:"hackerrank_r_dashboard.js"`
			HackerrankRDashboardCSSMap                                                                            string `json:"hackerrank_r_dashboard.css.map"`
			HackerrankRDashboardJsMap                                                                             string `json:"hackerrank_r_dashboard.js.map"`
			HackerrankRForumCSS                                                                                   string `json:"hackerrank_r_forum.css"`
			HackerrankRForumJs                                                                                    string `json:"hackerrank_r_forum.js"`
			HackerrankRForumCSSMap                                                                                string `json:"hackerrank_r_forum.css.map"`
			HackerrankRForumJsMap                                                                                 string `json:"hackerrank_r_forum.js.map"`
			HackerrankRInterviewCSS                                                                               string `json:"hackerrank_r_interview.css"`
			HackerrankRInterviewJs                                                                                string `json:"hackerrank_r_interview.js"`
			HackerrankRInterviewCSSMap                                                                            string `json:"hackerrank_r_interview.css.map"`
			HackerrankRInterviewJsMap                                                                             string `json:"hackerrank_r_interview.js.map"`
			HackerrankRJobsCSS                                                                                    string `json:"hackerrank_r_jobs.css"`
			HackerrankRJobsJs                                                                                     string `json:"hackerrank_r_jobs.js"`
			HackerrankRJobsCSSMap                                                                                 string `json:"hackerrank_r_jobs.css.map"`
			HackerrankRJobsJsMap                                                                                  string `json:"hackerrank_r_jobs.js.map"`
			HackerrankRKrackjackJs                                                                                string `json:"hackerrank_r_krackjack.js"`
			HackerrankRKrackjackJsMap                                                                             string `json:"hackerrank_r_krackjack.js.map"`
			HackerrankRKrackjackAdapterCSS                                                                        string `json:"hackerrank_r_krackjack_adapter.css"`
			HackerrankRKrackjackAdapterJs                                                                         string `json:"hackerrank_r_krackjack_adapter.js"`
			HackerrankRKrackjackAdapterCSSMap                                                                     string `json:"hackerrank_r_krackjack_adapter.css.map"`
			HackerrankRKrackjackAdapterJsMap                                                                      string `json:"hackerrank_r_krackjack_adapter.js.map"`
			HackerrankRLeaderboardCSS                                                                             string `json:"hackerrank_r_leaderboard.css"`
			HackerrankRLeaderboardJs                                                                              string `json:"hackerrank_r_leaderboard.js"`
			HackerrankRLeaderboardCSSMap                                                                          string `json:"hackerrank_r_leaderboard.css.map"`
			HackerrankRLeaderboardJsMap                                                                           string `json:"hackerrank_r_leaderboard.js.map"`
			HackerrankRLeaderboardv2CSS                                                                           string `json:"hackerrank_r_leaderboardv2.css"`
			HackerrankRLeaderboardv2Js                                                                            string `json:"hackerrank_r_leaderboardv2.js"`
			HackerrankRLeaderboardv2CSSMap                                                                        string `json:"hackerrank_r_leaderboardv2.css.map"`
			HackerrankRLeaderboardv2JsMap                                                                         string `json:"hackerrank_r_leaderboardv2.js.map"`
			HackerrankRLibraryCSS                                                                                 string `json:"hackerrank_r_library.css"`
			HackerrankRLibraryJs                                                                                  string `json:"hackerrank_r_library.js"`
			HackerrankRLibraryCSSMap                                                                              string `json:"hackerrank_r_library.css.map"`
			HackerrankRLibraryJsMap                                                                               string `json:"hackerrank_r_library.js.map"`
			HackerrankROldTrimmedCSS                                                                              string `json:"hackerrank_r_old_trimmed.css"`
			HackerrankROldTrimmedJs                                                                               string `json:"hackerrank_r_old_trimmed.js"`
			HackerrankROldTrimmedCSSMap                                                                           string `json:"hackerrank_r_old_trimmed.css.map"`
			HackerrankROldTrimmedJsMap                                                                            string `json:"hackerrank_r_old_trimmed.js.map"`
			HackerrankROnboardingCSS                                                                              string `json:"hackerrank_r_onboarding.css"`
			HackerrankROnboardingJs                                                                               string `json:"hackerrank_r_onboarding.js"`
			HackerrankROnboardingCSSMap                                                                           string `json:"hackerrank_r_onboarding.css.map"`
			HackerrankROnboardingJsMap                                                                            string `json:"hackerrank_r_onboarding.js.map"`
			HackerrankRPolyfillsIe11Js                                                                            string `json:"hackerrank_r_polyfills_ie11.js"`
			HackerrankRPolyfillsIe11JsMap                                                                         string `json:"hackerrank_r_polyfills_ie11.js.map"`
			HackerrankRProfileV2CSS                                                                               string `json:"hackerrank_r_profile_v2.css"`
			HackerrankRProfileV2Js                                                                                string `json:"hackerrank_r_profile_v2.js"`
			HackerrankRProfileV2CSSMap                                                                            string `json:"hackerrank_r_profile_v2.css.map"`
			HackerrankRProfileV2JsMap                                                                             string `json:"hackerrank_r_profile_v2.js.map"`
			HackerrankRRbaCSS                                                                                     string `json:"hackerrank_r_rba.css"`
			HackerrankRRbaJs                                                                                      string `json:"hackerrank_r_rba.js"`
			HackerrankRRbaCSSMap                                                                                  string `json:"hackerrank_r_rba.css.map"`
			HackerrankRRbaJsMap                                                                                   string `json:"hackerrank_r_rba.js.map"`
			HackerrankRScoringCSS                                                                                 string `json:"hackerrank_r_scoring.css"`
			HackerrankRScoringJs                                                                                  string `json:"hackerrank_r_scoring.js"`
			HackerrankRScoringCSSMap                                                                              string `json:"hackerrank_r_scoring.css.map"`
			HackerrankRScoringJsMap                                                                               string `json:"hackerrank_r_scoring.js.map"`
			HackerrankRSingletestCSS                                                                              string `json:"hackerrank_r_singletest.css"`
			HackerrankRSingletestJs                                                                               string `json:"hackerrank_r_singletest.js"`
			HackerrankRSingletestCSSMap                                                                           string `json:"hackerrank_r_singletest.css.map"`
			HackerrankRSingletestJsMap                                                                            string `json:"hackerrank_r_singletest.js.map"`
			HackerrankRSourcingApplicationsCSS                                                                    string `json:"hackerrank_r_sourcing_applications.css"`
			HackerrankRSourcingApplicationsJs                                                                     string `json:"hackerrank_r_sourcing_applications.js"`
			HackerrankRSourcingApplicationsCSSMap                                                                 string `json:"hackerrank_r_sourcing_applications.css.map"`
			HackerrankRSourcingApplicationsJsMap                                                                  string `json:"hackerrank_r_sourcing_applications.js.map"`
			HackerrankRSourcingCompanyCSS                                                                         string `json:"hackerrank_r_sourcing_company.css"`
			HackerrankRSourcingCompanyJs                                                                          string `json:"hackerrank_r_sourcing_company.js"`
			HackerrankRSourcingCompanyCSSMap                                                                      string `json:"hackerrank_r_sourcing_company.css.map"`
			HackerrankRSourcingCompanyJsMap                                                                       string `json:"hackerrank_r_sourcing_company.js.map"`
			HackerrankRSourcingCompanyJobsCSS                                                                     string `json:"hackerrank_r_sourcing_company_jobs.css"`
			HackerrankRSourcingCompanyJobsJs                                                                      string `json:"hackerrank_r_sourcing_company_jobs.js"`
			HackerrankRSourcingCompanyJobsCSSMap                                                                  string `json:"hackerrank_r_sourcing_company_jobs.css.map"`
			HackerrankRSourcingCompanyJobsJsMap                                                                   string `json:"hackerrank_r_sourcing_company_jobs.js.map"`
			HackerrankRSourcingMessagesCSS                                                                        string `json:"hackerrank_r_sourcing_messages.css"`
			HackerrankRSourcingMessagesJs                                                                         string `json:"hackerrank_r_sourcing_messages.js"`
			HackerrankRSourcingMessagesCSSMap                                                                     string `json:"hackerrank_r_sourcing_messages.css.map"`
			HackerrankRSourcingMessagesJsMap                                                                      string `json:"hackerrank_r_sourcing_messages.js.map"`
			HackerrankRSourcingSettingsCSS                                                                        string `json:"hackerrank_r_sourcing_settings.css"`
			HackerrankRSourcingSettingsJs                                                                         string `json:"hackerrank_r_sourcing_settings.js"`
			HackerrankRSourcingSettingsCSSMap                                                                     string `json:"hackerrank_r_sourcing_settings.css.map"`
			HackerrankRSourcingSettingsJsMap                                                                      string `json:"hackerrank_r_sourcing_settings.js.map"`
			HackerrankRTeamsCSS                                                                                   string `json:"hackerrank_r_teams.css"`
			HackerrankRTeamsJs                                                                                    string `json:"hackerrank_r_teams.js"`
			HackerrankRTeamsCSSMap                                                                                string `json:"hackerrank_r_teams.css.map"`
			HackerrankRTeamsJsMap                                                                                 string `json:"hackerrank_r_teams.js.map"`
			HackerrankRTestinsightsCSS                                                                            string `json:"hackerrank_r_testinsights.css"`
			HackerrankRTestinsightsJs                                                                             string `json:"hackerrank_r_testinsights.js"`
			HackerrankRTestinsightsCSSMap                                                                         string `json:"hackerrank_r_testinsights.css.map"`
			HackerrankRTestinsightsJsMap                                                                          string `json:"hackerrank_r_testinsights.js.map"`
			HackerrankRTestinviteCSS                                                                              string `json:"hackerrank_r_testinvite.css"`
			HackerrankRTestinviteJs                                                                               string `json:"hackerrank_r_testinvite.js"`
			HackerrankRTestinviteCSSMap                                                                           string `json:"hackerrank_r_testinvite.css.map"`
			HackerrankRTestinviteJsMap                                                                            string `json:"hackerrank_r_testinvite.js.map"`
			HackerrankRTestoverviewCSS                                                                            string `json:"hackerrank_r_testoverview.css"`
			HackerrankRTestoverviewJs                                                                             string `json:"hackerrank_r_testoverview.js"`
			HackerrankRTestoverviewCSSMap                                                                         string `json:"hackerrank_r_testoverview.css.map"`
			HackerrankRTestoverviewJsMap                                                                          string `json:"hackerrank_r_testoverview.js.map"`
			HackerrankRTestquestionsCSS                                                                           string `json:"hackerrank_r_testquestions.css"`
			HackerrankRTestquestionsJs                                                                            string `json:"hackerrank_r_testquestions.js"`
			HackerrankRTestquestionsCSSMap                                                                        string `json:"hackerrank_r_testquestions.css.map"`
			HackerrankRTestquestionsJsMap                                                                         string `json:"hackerrank_r_testquestions.js.map"`
			HackerrankRTestreportslistingCSS                                                                      string `json:"hackerrank_r_testreportslisting.css"`
			HackerrankRTestreportslistingJs                                                                       string `json:"hackerrank_r_testreportslisting.js"`
			HackerrankRTestreportslistingCSSMap                                                                   string `json:"hackerrank_r_testreportslisting.css.map"`
			HackerrankRTestreportslistingJsMap                                                                    string `json:"hackerrank_r_testreportslisting.js.map"`
			HackerrankRTestsettingsCSS                                                                            string `json:"hackerrank_r_testsettings.css"`
			HackerrankRTestsettingsJs                                                                             string `json:"hackerrank_r_testsettings.js"`
			HackerrankRTestsettingsCSSMap                                                                         string `json:"hackerrank_r_testsettings.css.map"`
			HackerrankRTestsettingsJsMap                                                                          string `json:"hackerrank_r_testsettings.js.map"`
			HackerrankRWorkCSS                                                                                    string `json:"hackerrank_r_work.css"`
			HackerrankRWorkJs                                                                                     string `json:"hackerrank_r_work.js"`
			HackerrankRWorkCSSMap                                                                                 string `json:"hackerrank_r_work.css.map"`
			HackerrankRWorkJsMap                                                                                  string `json:"hackerrank_r_work.js.map"`
			UIIconsFontFaceCSS                                                                                    string `json:"ui_icons_font_face.css"`
			UIIconsFontFaceJs                                                                                     string `json:"ui_icons_font_face.js"`
			UIIconsFontFaceCSSMap                                                                                 string `json:"ui_icons_font_face.css.map"`
			UIIconsFontFaceJsMap                                                                                  string `json:"ui_icons_font_face.js.map"`
			VendorsHackerrankRCalendarJs                                                                          string `json:"vendors~hackerrank_r_calendar.js"`
			VendorsHackerrankRCalendarJsMap                                                                       string `json:"vendors~hackerrank_r_calendar.js.map"`
			VendorsHackerrankRChartsCSS                                                                           string `json:"vendors~hackerrank_r_charts.css"`
			VendorsHackerrankRChartsJs                                                                            string `json:"vendors~hackerrank_r_charts.js"`
			VendorsHackerrankRChartsCSSMap                                                                        string `json:"vendors~hackerrank_r_charts.css.map"`
			VendorsHackerrankRChartsJsMap                                                                         string `json:"vendors~hackerrank_r_charts.js.map"`
			VendorsHackerrankRCodeshellLibJs                                                                      string `json:"vendors~hackerrank_r_codeshell_lib.js"`
			VendorsHackerrankRCodeshellLibJsMap                                                                   string `json:"vendors~hackerrank_r_codeshell_lib.js.map"`
			VendorsHackerrankRIntersectionObserverJs                                                              string `json:"vendors~hackerrank_r_intersection_observer.js"`
			VendorsHackerrankRIntersectionObserverJsMap                                                           string `json:"vendors~hackerrank_r_intersection_observer.js.map"`
			VendorsHackerrankRInterviewCSS                                                                        string `json:"vendors~hackerrank_r_interview.css"`
			VendorsHackerrankRInterviewJs                                                                         string `json:"vendors~hackerrank_r_interview.js"`
			VendorsHackerrankRInterviewCSSMap                                                                     string `json:"vendors~hackerrank_r_interview.css.map"`
			VendorsHackerrankRInterviewJsMap                                                                      string `json:"vendors~hackerrank_r_interview.js.map"`
			VendorsHackerrankRJobsCSS                                                                             string `json:"vendors~hackerrank_r_jobs.css"`
			VendorsHackerrankRJobsJs                                                                              string `json:"vendors~hackerrank_r_jobs.js"`
			VendorsHackerrankRJobsCSSMap                                                                          string `json:"vendors~hackerrank_r_jobs.css.map"`
			VendorsHackerrankRJobsJsMap                                                                           string `json:"vendors~hackerrank_r_jobs.js.map"`
			VendorsHackerrankRKrackjackCSS                                                                        string `json:"vendors~hackerrank_r_krackjack.css"`
			VendorsHackerrankRKrackjackJs                                                                         string `json:"vendors~hackerrank_r_krackjack.js"`
			VendorsHackerrankRKrackjackCSSMap                                                                     string `json:"vendors~hackerrank_r_krackjack.css.map"`
			VendorsHackerrankRKrackjackJsMap                                                                      string `json:"vendors~hackerrank_r_krackjack.js.map"`
			VendorsHackerrankRLeaderboardCSS                                                                      string `json:"vendors~hackerrank_r_leaderboard.css"`
			VendorsHackerrankRLeaderboardJs                                                                       string `json:"vendors~hackerrank_r_leaderboard.js"`
			VendorsHackerrankRLeaderboardCSSMap                                                                   string `json:"vendors~hackerrank_r_leaderboard.css.map"`
			VendorsHackerrankRLeaderboardJsMap                                                                    string `json:"vendors~hackerrank_r_leaderboard.js.map"`
			VendorsHackerrankRSourcingCompanyJobsCSS                                                              string `json:"vendors~hackerrank_r_sourcing_company_jobs.css"`
			VendorsHackerrankRSourcingCompanyJobsJs                                                               string `json:"vendors~hackerrank_r_sourcing_company_jobs.js"`
			VendorsHackerrankRSourcingCompanyJobsCSSMap                                                           string `json:"vendors~hackerrank_r_sourcing_company_jobs.css.map"`
			VendorsHackerrankRSourcingCompanyJobsJsMap                                                            string `json:"vendors~hackerrank_r_sourcing_company_jobs.js.map"`
			VendorsHackerrankRTeamsCSS                                                                            string `json:"vendors~hackerrank_r_teams.css"`
			VendorsHackerrankRTeamsJs                                                                             string `json:"vendors~hackerrank_r_teams.js"`
			VendorsHackerrankRTeamsCSSMap                                                                         string `json:"vendors~hackerrank_r_teams.css.map"`
			VendorsHackerrankRTeamsJsMap                                                                          string `json:"vendors~hackerrank_r_teams.js.map"`
			VendorsHackerrankRTestinsightsJs                                                                      string `json:"vendors~hackerrank_r_testinsights.js"`
			VendorsHackerrankRTestinsightsJsMap                                                                   string `json:"vendors~hackerrank_r_testinsights.js.map"`
			VendorsHackerrankRTestoverviewCSS                                                                     string `json:"vendors~hackerrank_r_testoverview.css"`
			VendorsHackerrankRTestoverviewJs                                                                      string `json:"vendors~hackerrank_r_testoverview.js"`
			VendorsHackerrankRTestoverviewCSSMap                                                                  string `json:"vendors~hackerrank_r_testoverview.css.map"`
			VendorsHackerrankRTestoverviewJsMap                                                                   string `json:"vendors~hackerrank_r_testoverview.js.map"`
			VendorsHackerrankRTestquestionsJs                                                                     string `json:"vendors~hackerrank_r_testquestions.js"`
			VendorsHackerrankRTestquestionsJsMap                                                                  string `json:"vendors~hackerrank_r_testquestions.js.map"`
			VendorsHackerrankRTestreportslistingCSS                                                               string `json:"vendors~hackerrank_r_testreportslisting.css"`
			VendorsHackerrankRTestreportslistingJs                                                                string `json:"vendors~hackerrank_r_testreportslisting.js"`
			VendorsHackerrankRTestreportslistingCSSMap                                                            string `json:"vendors~hackerrank_r_testreportslisting.css.map"`
			VendorsHackerrankRTestreportslistingJsMap                                                             string `json:"vendors~hackerrank_r_testreportslisting.js.map"`
			VendorsHackerrankRTestsettingsCSS                                                                     string `json:"vendors~hackerrank_r_testsettings.css"`
			VendorsHackerrankRTestsettingsJs                                                                      string `json:"vendors~hackerrank_r_testsettings.js"`
			VendorsHackerrankRTestsettingsCSSMap                                                                  string `json:"vendors~hackerrank_r_testsettings.css.map"`
			VendorsHackerrankRTestsettingsJsMap                                                                   string `json:"vendors~hackerrank_r_testsettings.js.map"`
			VendorsLspClientJs                                                                                    string `json:"vendors~lsp_client.js"`
			VendorsLspClientJsMap                                                                                 string `json:"vendors~lsp_client.js.map"`
			VendorsMonacoemacsJs                                                                                  string `json:"vendors~monacoemacs.js"`
			VendorsMonacoemacsJsMap                                                                               string `json:"vendors~monacoemacs.js.map"`
			VendorsMonacovimJs                                                                                    string `json:"vendors~monacovim.js"`
			VendorsMonacovimJsMap                                                                                 string `json:"vendors~monacovim.js.map"`
			One0886C8Cd8C06Dbbe6037C9Js                                                                           string `json:"108-86c8cd8c06dbbe6037c9.js"`
			Sourcemaps10886C8Cd8C06Dbbe6037C9JsMap                                                                string `json:"sourcemaps/108-86c8cd8c06dbbe6037c9.js.map"`
			One099244Cd37C78C7519E310Js                                                                           string `json:"109-9244cd37c78c7519e310.js"`
			Sourcemaps1099244Cd37C78C7519E310JsMap                                                                string `json:"sourcemaps/109-9244cd37c78c7519e310.js.map"`
			One10531212408D25654F4Ff3Js                                                                           string `json:"110-531212408d25654f4ff3.js"`
			Sourcemaps110531212408D25654F4Ff3JsMap                                                                string `json:"sourcemaps/110-531212408d25654f4ff3.js.map"`
			One11Df107A7F6D60A76E3946Js                                                                           string `json:"111-df107a7f6d60a76e3946.js"`
			Sourcemaps111Df107A7F6D60A76E3946JsMap                                                                string `json:"sourcemaps/111-df107a7f6d60a76e3946.js.map"`
			One12781580Dcafd73Dbaf2F4Js                                                                           string `json:"112-781580dcafd73dbaf2f4.js"`
			Sourcemaps112781580Dcafd73Dbaf2F4JsMap                                                                string `json:"sourcemaps/112-781580dcafd73dbaf2f4.js.map"`
			One1329967368Cbbe84Cba3C0Js                                                                           string `json:"113-29967368cbbe84cba3c0.js"`
			Sourcemaps11329967368Cbbe84Cba3C0JsMap                                                                string `json:"sourcemaps/113-29967368cbbe84cba3c0.js.map"`
			One1403F28Bcd9Cc392Df44CbJs                                                                           string `json:"114-03f28bcd9cc392df44cb.js"`
			Sourcemaps11403F28Bcd9Cc392Df44CbJsMap                                                                string `json:"sourcemaps/114-03f28bcd9cc392df44cb.js.map"`
			One1580A0A463F7175210F3A4Js                                                                           string `json:"115-80a0a463f7175210f3a4.js"`
			Sourcemaps11580A0A463F7175210F3A4JsMap                                                                string `json:"sourcemaps/115-80a0a463f7175210f3a4.js.map"`
			One164393655A3F0183272F29Js                                                                           string `json:"116-4393655a3f0183272f29.js"`
			Sourcemaps1164393655A3F0183272F29JsMap                                                                string `json:"sourcemaps/116-4393655a3f0183272f29.js.map"`
			One17A7Cb3B37533E62581Ae6Js                                                                           string `json:"117-a7cb3b37533e62581ae6.js"`
			Sourcemaps117A7Cb3B37533E62581Ae6JsMap                                                                string `json:"sourcemaps/117-a7cb3b37533e62581ae6.js.map"`
			One184C567Bc156F0C3466F04Js                                                                           string `json:"118-4c567bc156f0c3466f04.js"`
			Sourcemaps1184C567Bc156F0C3466F04JsMap                                                                string `json:"sourcemaps/118-4c567bc156f0c3466f04.js.map"`
			One19Afcef62863B037B9F4CfJs                                                                           string `json:"119-afcef62863b037b9f4cf.js"`
			Sourcemaps119Afcef62863B037B9F4CfJsMap                                                                string `json:"sourcemaps/119-afcef62863b037b9f4cf.js.map"`
			One20Ce8Ceede3Ee8Ebaa4459Js                                                                           string `json:"120-ce8ceede3ee8ebaa4459.js"`
			Sourcemaps120Ce8Ceede3Ee8Ebaa4459JsMap                                                                string `json:"sourcemaps/120-ce8ceede3ee8ebaa4459.js.map"`
			One2174D711Dc9E37833077A3Js                                                                           string `json:"121-74d711dc9e37833077a3.js"`
			Sourcemaps12174D711Dc9E37833077A3JsMap                                                                string `json:"sourcemaps/121-74d711dc9e37833077a3.js.map"`
			One22302A1Fe587Fce19A6D98Js                                                                           string `json:"122-302a1fe587fce19a6d98.js"`
			Sourcemaps122302A1Fe587Fce19A6D98JsMap                                                                string `json:"sourcemaps/122-302a1fe587fce19a6d98.js.map"`
			One237C8B8Ad88530C15485E6Js                                                                           string `json:"123-7c8b8ad88530c15485e6.js"`
			Sourcemaps1237C8B8Ad88530C15485E6JsMap                                                                string `json:"sourcemaps/123-7c8b8ad88530c15485e6.js.map"`
			One24E07B0185C6923A1Dc9B8Js                                                                           string `json:"124-e07b0185c6923a1dc9b8.js"`
			Sourcemaps124E07B0185C6923A1Dc9B8JsMap                                                                string `json:"sourcemaps/124-e07b0185c6923a1dc9b8.js.map"`
			One251C7E55F536C14Dad958DJs                                                                           string `json:"125-1c7e55f536c14dad958d.js"`
			Sourcemaps1251C7E55F536C14Dad958DJsMap                                                                string `json:"sourcemaps/125-1c7e55f536c14dad958d.js.map"`
			One2608016Ea1616C077F5De0Js                                                                           string `json:"126-08016ea1616c077f5de0.js"`
			Sourcemaps12608016Ea1616C077F5De0JsMap                                                                string `json:"sourcemaps/126-08016ea1616c077f5de0.js.map"`
			One2793E661231416E0Db1243Js                                                                           string `json:"127-93e661231416e0db1243.js"`
			Sourcemaps12793E661231416E0Db1243JsMap                                                                string `json:"sourcemaps/127-93e661231416e0db1243.js.map"`
			One28Bb32C9C72Ce69F014A00Js                                                                           string `json:"128-bb32c9c72ce69f014a00.js"`
			Sourcemaps128Bb32C9C72Ce69F014A00JsMap                                                                string `json:"sourcemaps/128-bb32c9c72ce69f014a00.js.map"`
			One2945E2E14Ba099E66F0042Js                                                                           string `json:"129-45e2e14ba099e66f0042.js"`
			Sourcemaps12945E2E14Ba099E66F0042JsMap                                                                string `json:"sourcemaps/129-45e2e14ba099e66f0042.js.map"`
			One303A4E8A182Ceb5574Ca1DJs                                                                           string `json:"130-3a4e8a182ceb5574ca1d.js"`
			Sourcemaps1303A4E8A182Ceb5574Ca1DJsMap                                                                string `json:"sourcemaps/130-3a4e8a182ceb5574ca1d.js.map"`
			One31A37Ce0Cb459Eb23272A7Js                                                                           string `json:"131-a37ce0cb459eb23272a7.js"`
			Sourcemaps131A37Ce0Cb459Eb23272A7JsMap                                                                string `json:"sourcemaps/131-a37ce0cb459eb23272a7.js.map"`
			One32E50C32Fde3B26136B85EJs                                                                           string `json:"132-e50c32fde3b26136b85e.js"`
			Sourcemaps132E50C32Fde3B26136B85EJsMap                                                                string `json:"sourcemaps/132-e50c32fde3b26136b85e.js.map"`
			One3304A48F33D422Bcc3A782Js                                                                           string `json:"133-04a48f33d422bcc3a782.js"`
			Sourcemaps13304A48F33D422Bcc3A782JsMap                                                                string `json:"sourcemaps/133-04a48f33d422bcc3a782.js.map"`
			One34729F7551B68029A9B2D5Js                                                                           string `json:"134-729f7551b68029a9b2d5.js"`
			Sourcemaps134729F7551B68029A9B2D5JsMap                                                                string `json:"sourcemaps/134-729f7551b68029a9b2d5.js.map"`
			One35Ee61465Ff318E51998A5Js                                                                           string `json:"135-ee61465ff318e51998a5.js"`
			Sourcemaps135Ee61465Ff318E51998A5JsMap                                                                string `json:"sourcemaps/135-ee61465ff318e51998a5.js.map"`
			One361A5072D729Df7B741222Js                                                                           string `json:"136-1a5072d729df7b741222.js"`
			Sourcemaps1361A5072D729Df7B741222JsMap                                                                string `json:"sourcemaps/136-1a5072d729df7b741222.js.map"`
			One3714D322B72797Ad99FebdJs                                                                           string `json:"137-14d322b72797ad99febd.js"`
			Sourcemaps13714D322B72797Ad99FebdJsMap                                                                string `json:"sourcemaps/137-14d322b72797ad99febd.js.map"`
			One388Ca7Cd0B8468728B97AdJs                                                                           string `json:"138-8ca7cd0b8468728b97ad.js"`
			Sourcemaps1388Ca7Cd0B8468728B97AdJsMap                                                                string `json:"sourcemaps/138-8ca7cd0b8468728b97ad.js.map"`
			One39E734A910F53C1424C83AJs                                                                           string `json:"139-e734a910f53c1424c83a.js"`
			Sourcemaps139E734A910F53C1424C83AJsMap                                                                string `json:"sourcemaps/139-e734a910f53c1424c83a.js.map"`
			One40611Abd3Add32122C4A1AJs                                                                           string `json:"140-611abd3add32122c4a1a.js"`
			Sourcemaps140611Abd3Add32122C4A1AJsMap                                                                string `json:"sourcemaps/140-611abd3add32122c4a1a.js.map"`
			One413Ecba6E9A008A2Db85C4Js                                                                           string `json:"141-3ecba6e9a008a2db85c4.js"`
			Sourcemaps1413Ecba6E9A008A2Db85C4JsMap                                                                string `json:"sourcemaps/141-3ecba6e9a008a2db85c4.js.map"`
			One4250Ff538779D997197458Js                                                                           string `json:"142-50ff538779d997197458.js"`
			Sourcemaps14250Ff538779D997197458JsMap                                                                string `json:"sourcemaps/142-50ff538779d997197458.js.map"`
			One43709D9D1Ed3Eda989Da74Js                                                                           string `json:"143-709d9d1ed3eda989da74.js"`
			Sourcemaps143709D9D1Ed3Eda989Da74JsMap                                                                string `json:"sourcemaps/143-709d9d1ed3eda989da74.js.map"`
			One44745C939A137C234817D8Js                                                                           string `json:"144-745c939a137c234817d8.js"`
			Sourcemaps144745C939A137C234817D8JsMap                                                                string `json:"sourcemaps/144-745c939a137c234817d8.js.map"`
			One458Df3D15Cbe24824126C2Js                                                                           string `json:"145-8df3d15cbe24824126c2.js"`
			Sourcemaps1458Df3D15Cbe24824126C2JsMap                                                                string `json:"sourcemaps/145-8df3d15cbe24824126c2.js.map"`
			One46C4F7F4A5E54A4132Aa5FJs                                                                           string `json:"146-c4f7f4a5e54a4132aa5f.js"`
			Sourcemaps146C4F7F4A5E54A4132Aa5FJsMap                                                                string `json:"sourcemaps/146-c4f7f4a5e54a4132aa5f.js.map"`
			One47063Bbea337Afe704Fb96Js                                                                           string `json:"147-063bbea337afe704fb96.js"`
			Sourcemaps147063Bbea337Afe704Fb96JsMap                                                                string `json:"sourcemaps/147-063bbea337afe704fb96.js.map"`
			One488C2B28B2Bb27A73Be0A3Js                                                                           string `json:"148-8c2b28b2bb27a73be0a3.js"`
			Sourcemaps1488C2B28B2Bb27A73Be0A3JsMap                                                                string `json:"sourcemaps/148-8c2b28b2bb27a73be0a3.js.map"`
			One49060E80Bbc5Db710C0Ff3Js                                                                           string `json:"149-060e80bbc5db710c0ff3.js"`
			Sourcemaps149060E80Bbc5Db710C0Ff3JsMap                                                                string `json:"sourcemaps/149-060e80bbc5db710c0ff3.js.map"`
			One503C93Fcf3Bde4E61Dfea4Js                                                                           string `json:"150-3c93fcf3bde4e61dfea4.js"`
			Sourcemaps1503C93Fcf3Bde4E61Dfea4JsMap                                                                string `json:"sourcemaps/150-3c93fcf3bde4e61dfea4.js.map"`
			One512436Ef7D1604897Cf6D3Js                                                                           string `json:"151-2436ef7d1604897cf6d3.js"`
			Sourcemaps1512436Ef7D1604897Cf6D3JsMap                                                                string `json:"sourcemaps/151-2436ef7d1604897cf6d3.js.map"`
			One521C5E8A5Af97B8A98869AJs                                                                           string `json:"152-1c5e8a5af97b8a98869a.js"`
			Sourcemaps1521C5E8A5Af97B8A98869AJsMap                                                                string `json:"sourcemaps/152-1c5e8a5af97b8a98869a.js.map"`
			One536832Ae4D4393A688B684Js                                                                           string `json:"153-6832ae4d4393a688b684.js"`
			Sourcemaps1536832Ae4D4393A688B684JsMap                                                                string `json:"sourcemaps/153-6832ae4d4393a688b684.js.map"`
			One540786D8Dfe358F411B2A6Js                                                                           string `json:"154-0786d8dfe358f411b2a6.js"`
			Sourcemaps1540786D8Dfe358F411B2A6JsMap                                                                string `json:"sourcemaps/154-0786d8dfe358f411b2a6.js.map"`
			One55Da3F8647Ef6Cb87F0F0EJs                                                                           string `json:"155-da3f8647ef6cb87f0f0e.js"`
			Sourcemaps155Da3F8647Ef6Cb87F0F0EJsMap                                                                string `json:"sourcemaps/155-da3f8647ef6cb87f0f0e.js.map"`
			One56D4Af299207E09Ca88349Js                                                                           string `json:"156-d4af299207e09ca88349.js"`
			Sourcemaps156D4Af299207E09Ca88349JsMap                                                                string `json:"sourcemaps/156-d4af299207e09ca88349.js.map"`
			One5710E7025B1376E753913DJs                                                                           string `json:"157-10e7025b1376e753913d.js"`
			Sourcemaps15710E7025B1376E753913DJsMap                                                                string `json:"sourcemaps/157-10e7025b1376e753913d.js.map"`
			One58D5816C9Bba64962C8347Js                                                                           string `json:"158-d5816c9bba64962c8347.js"`
			Sourcemaps158D5816C9Bba64962C8347JsMap                                                                string `json:"sourcemaps/158-d5816c9bba64962c8347.js.map"`
			One59A23960A69E09Dd6C0B18Js                                                                           string `json:"159-a23960a69e09dd6c0b18.js"`
			Sourcemaps159A23960A69E09Dd6C0B18JsMap                                                                string `json:"sourcemaps/159-a23960a69e09dd6c0b18.js.map"`
			CSSWorkerJs                                                                                           string `json:"css.worker.js"`
			EditorWorkerJs                                                                                        string `json:"editor.worker.js"`
			HTMLWorkerJs                                                                                          string `json:"html.worker.js"`
			JSONWorkerJs                                                                                          string `json:"json.worker.js"`
			SourcemapsCSSWorkerJsMap                                                                              string `json:"sourcemaps/css.worker.js.map"`
			SourcemapsEditorWorkerJsMap                                                                           string `json:"sourcemaps/editor.worker.js.map"`
			SourcemapsHTMLWorkerJsMap                                                                             string `json:"sourcemaps/html.worker.js.map"`
			SourcemapsJSONWorkerJsMap                                                                             string `json:"sourcemaps/json.worker.js.map"`
			SourcemapsTypescriptWorkerJsMap                                                                       string `json:"sourcemaps/typescript.worker.js.map"`
			TypescriptWorkerJs                                                                                    string `json:"typescript.worker.js"`
		} `json:"manifest"`
		Bundles struct {
			HackerrankRCommunityJs bool `json:"hackerrank_r_community.js"`
			HackerrankRAppJs       bool `json:"hackerrank_r_app.js"`
			HackerrankRCommonJs    bool `json:"hackerrank_r_common.js"`
			HackerrankRChallengeJs bool `json:"hackerrank_r_challenge.js"`
		} `json:"bundles"`
		Slugs struct {
			Master struct {
				Type string `json:"type"`
			} `json:"master"`
		} `json:"slugs"`
		ProductNamespace        string      `json:"productNamespace"`
		IsMobile                bool        `json:"isMobile"`
		Host                    string      `json:"host"`
		ENV                     string      `json:"ENV"`
		AssetPath               string      `json:"asset_path"`
		AssetHost               string      `json:"asset_host"`
		AssetHostPath           string      `json:"asset_host_path"`
		FilepickerKey           string      `json:"filepicker_key"`
		PubsubHost              string      `json:"pubsub_host"`
		ReleaseVersion          int         `json:"release_version"`
		LandingContestSlug      string      `json:"landing_contest_slug"`
		CurrentContestNamespace interface{} `json:"current_contest_namespace"`
		UsingContestNamespace   interface{} `json:"using_contest_namespace"`
		CsrfToken               string      `json:"csrfToken"`
		CookieHeaders           []string    `json:"cookieHeaders"`
		AdminSu                 interface{} `json:"adminSu"`
		AssetPathHost           string      `json:"asset_path_host"`
	} `json:"metadata"`
	Abtest struct {
		Models []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Product  int    `json:"product"`
			Variants []struct {
				ID      int    `json:"id"`
				Weight  int    `json:"weight"`
				Variant string `json:"variant"`
			} `json:"variants"`
			HackerVariant struct {
				Variant string `json:"variant"`
				Forced  bool   `json:"forced"`
			} `json:"hacker_variant"`
		} `json:"models"`
	} `json:"abtest"`
	App struct {
		ClientInitialLoading bool          `json:"clientInitialLoading"`
		ToastMessage         interface{}   `json:"toastMessage"`
		Countries            []interface{} `json:"countries"`
	} `json:"app"`
	FeatureFeedback struct {
		FeatureLists struct {
		} `json:"featureLists"`
	} `json:"featureFeedback"`
}
