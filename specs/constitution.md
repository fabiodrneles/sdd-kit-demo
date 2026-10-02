# Constituição do sdd-kit-demo

Princípios que toda spec e todo PR devem respeitar. Mudá-los exige uma spec própria.

1. **Testado.** Todo critério de aceite tem teste automatizado; o CI bloqueia merge vermelho.
2. **Falhar alto.** Link quebrado gera código de saída diferente de zero e uma linha com arquivo, linha e motivo.
3. **Determinístico por padrão.** Sem `--external`, nenhuma verificação depende de rede.
4. **Saída para pessoas e máquinas.** Toda informação do texto também existe no formato JSON.
5. **Sem dependências desnecessárias.** Só a biblioteca padrão do Go, salvo decisão registrada.
6. **Documentação verificada.** Cada comando do README roda no CI.
